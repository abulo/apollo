package monitor

import (
	"context"
	"fmt"
	"io"
	httpnet "net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/abulo/ratel/v3/core/env"
	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/guonaihong/gout"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GetLocalIP 获取服务器内网IP
func GetLocalIP() (ip string, err error) {
	addrs, err := httpnet.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*httpnet.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

const (
	whoisApi = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip="
	dyndns   = "http://members.3322.org/dyndns/getip" // 备用："https://ifconfig.co/ip"
)

type WhoisRegionData struct {
	Ip         string `json:"ip,omitempty"`
	Pro        string `json:"pro,omitempty"`
	ProCode    string `json:"proCode,omitempty"`
	City       string `json:"city,omitempty"`
	CityCode   string `json:"cityCode,omitempty"`
	Region     string `json:"region,omitempty"`
	RegionCode string `json:"regionCode,omitempty"`
	Addr       string `json:"addr,omitempty"`
	Err        string `json:"err,omitempty"`
}

// GetPublicIP 获取公网IP
func GetPublicIP() (ip string, err error) {
	var data *WhoisRegionData
	err = gout.GET(whoisApi).BindJSON(data).Do()
	if err != nil {
		return GetPublicIP2()
	}
	if data == nil {
		return "0.0.0.0", nil
	}
	return data.Ip, nil
}

func GetPublicIP2() (ip string, err error) {
	response, err := http.Get(dyndns)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	ip = strings.ReplaceAll(string(body), "\n", "")
	return
}

type Monitor struct {
	// LaunchTime 启动时间
	LaunchTime null.Time `json:"launchTime"`
	// RunTime 运行时间
	RunTime int64 `json:"runTime"`
	// BuildTime 编译时间
	BuildTime null.Time `json:"buildTime"`
	// BuildVersion 编译版本
	BuildVersion string `json:"buildVersion"`
	// RatelVersion Ratel版本
	RatelVersion string `json:"ratelVersion"`
	// GoVersion Go版本
	GoVersion string `json:"goVersion"`
	// GoArch Go架构
	GoArch string `json:"goArch"`
	// 内网IP
	IntranetIP string `json:"intranetIP"`
	// 公网IP
	PublicIP string `json:"publicIP"`
	// NetIO 网络流量统计
	NetIO []*NetIO `json:"netIO"`
	// LoadAvg 负载统计
	LoadAvg []*LoadAvg `json:"loadAvg"`
}

type NetIO struct {
	Time      null.Time `json:"time"`
	BytesSent uint64    `json:"bytesSent"` // number of bytes sent
	BytesRecv uint64    `json:"bytesRecv"` // number of bytes received
	Down      float64   `json:"down"`
	UP        float64   `json:"up"`
}

type LoadAvg struct {
	Time  null.Time `json:"time"`
	Avg   float64   `json:"avg"`
	Ratio float64   `json:"ratio"`
}

type MonitorHeadExtra struct {
	Data  interface{} `json:"data"`
	Data1 string      `json:"data1,omitempty"`
}
type MonitorHead struct {
	Title       string           `json:"title"`
	Data        string           `json:"data"`
	BottomTitle string           `json:"bottomTitle"`
	TotalSum    string           `json:"totalSum"`
	IconClass   string           `json:"iconClass"`
	Extra       MonitorHeadExtra `json:"extra"`
}

var MonitorCore *Monitor

func getMonitor() *Monitor {
	if MonitorCore == nil {
		MonitorCore = &Monitor{}
		intranetIP, _ := GetLocalIP()
		publicIP, _ := GetPublicIP()
		MonitorCore.IntranetIP = intranetIP
		MonitorCore.PublicIP = publicIP
		MonitorCore.GoVersion = runtime.Version()
		MonitorCore.GoArch = runtime.GOARCH
		MonitorCore.BuildVersion = env.BuildVersion()
		MonitorCore.RatelVersion = env.RatelVersion()
		MonitorCore.BuildTime = null.TimeFrom(cast.ToTimeInDefaultLocation(env.BuildTime(), time.Local))
		MonitorCore.LaunchTime = null.TimeFrom(cast.ToTimeInDefaultLocation(env.StartTime(), time.Local))
	}
	MonitorCore.RunTime = util.Now().Unix() - cast.ToTimeInDefaultLocation(env.StartTime(), time.Local).Unix()
	MonitorCore.loadAvg()
	MonitorCore.netIO()
	return MonitorCore
}

func (s *Monitor) loadAvg() {
	pl, _ := load.Avg()
	counter := LoadAvg{
		Time:  null.TimeFrom(util.Now()),
		Avg:   pl.Load1,
		Ratio: pl.Load1 / (float64(runtime.NumCPU()) * 2) * 100,
	}
	s.LoadAvg = append(s.LoadAvg, &counter)
	if len(s.LoadAvg) > 10 {
		s.LoadAvg = append(s.LoadAvg[:0], s.LoadAvg[(1):]...)
	}
}
func (s *Monitor) netIO() {
	var counter NetIO
	ni, _ := net.IOCounters(true)
	counter.Time = null.TimeFrom(util.Now())
	for _, v := range ni {
		counter.BytesSent += v.BytesSent
		counter.BytesRecv += v.BytesRecv
	}

	if len(s.NetIO) > 0 {
		lastNetIO := s.NetIO[len(s.NetIO)-1]
		sub := counter.Time.Time.Sub(lastNetIO.Time.Time).Seconds()
		counter.Down = Round2Float64((float64(counter.BytesRecv - lastNetIO.BytesRecv)) / sub)
		counter.UP = Round2Float64((float64(counter.BytesSent - lastNetIO.BytesSent)) / sub)
	}

	s.NetIO = append(s.NetIO, &counter)
	if len(s.NetIO) > 10 {
		s.NetIO = append(s.NetIO[:0], s.NetIO[(1):]...)
	}
}

// Round2String 四舍五入保留小数，默认2位
func Round2String(value float64, args ...interface{}) string {
	var places = 2
	if len(args) > 0 {
		places = cast.ToInt(args[0])
	}
	return fmt.Sprintf("%0."+strconv.Itoa(places)+"f", value)
}

// Round2Float64 四舍五入保留小数，默认2位
func Round2Float64(value float64, args ...interface{}) float64 {
	return cast.ToFloat64(Round2String(value, args...))
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(data int64) string {
	var factor float64 = 1024
	res := float64(data)
	for _, unit := range []string{"", "K", "M", "G", "T", "P"} {
		if res < factor {
			return fmt.Sprintf("%.2f%sB", res, unit)
		}
		res /= factor
	}
	return fmt.Sprintf("%.2f%sB", res, "P")
}

// 文件信息
type fileInfo struct {
	name string
	size int64
}

// WalkDir 递归获取目录下文件的名称和大小
func WalkDir(dirname string) (error, []fileInfo) {
	op, err := filepath.Abs(dirname) // 获取目录的绝对路径
	if nil != err {
		return err, nil
	}
	files, err := os.ReadDir(op) // 获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		return err, nil
	}

	var fileInfos []fileInfo // 返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			err, fs := WalkDir(op + `/` + f.Name()) // 路径分隔符，linux 和 windows 不同
			if nil != err {
				return err, nil
			}
			fileInfos = append(fileInfos, fs...) // 将 slice 添加到 slice
		} else {
			info, err := f.Info()
			if nil != err {
				return err, nil
			}
			fi := fileInfo{op + `/` + f.Name(), info.Size()}
			fileInfos = append(fileInfos, fi) // slice 中添加成员
		}
	}
	return nil, fileInfos
}

func DirSize(dirname string) string {
	var (
		ss       int64
		_, files = WalkDir(dirname)
	)
	for _, n := range files {
		ss += n.size
	}
	return FileSize(ss)
}

func RunInfo(ctx context.Context, newCtx *app.RequestContext) {
	// 客户端可以通过 Last-Event-ID 告知服务器收到的最后一个事件
	// lastEventID := sse.GetLastEventID(newCtx)

	var (
		meta     = getMonitor()
		mHost, _ = host.Info()
		pwd, _   = os.Getwd()
		gm       runtime.MemStats
	)
	runtime.ReadMemStats(&gm)
	data := map[string]interface{}{
		// 服务器信息
		"hostname":     mHost.Hostname,
		"os":           mHost.OS,
		"arch":         mHost.KernelArch,
		"platform":     mHost.Platform,
		"launchTime":   meta.LaunchTime,
		"runTime":      meta.RunTime,
		"buildTime":    meta.BuildTime,
		"buildVersion": meta.BuildVersion,
		"ratelVersion": meta.RatelVersion,
		"goVersion":    meta.GoVersion,
		"goArch":       meta.GoArch,
		"intranetIP":   meta.IntranetIP,
		"publicIP":     meta.PublicIP,
		"rootPath":     runtime.GOROOT(),
		"pwd":          pwd,
		"goroutine":    runtime.NumGoroutine(),
		"goMem":        FileSize(int64(gm.Sys)),
		"goSize":       DirSize(pwd),
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": data,
	})
}

func Trend(ctx context.Context, newCtx *app.RequestContext) {
	type NetC struct {
		Time      null.Time `json:"time"`
		BytesSent string    `json:"bytesSent"` // number of bytes sent
		BytesRecv string    `json:"bytesRecv"` // number of bytes received
		Down      float64   `json:"down"`
		UP        float64   `json:"up"`
	}

	var (
		mCpu, cpuErr         = cpu.Info()
		mCpuUsed             float64
		mMem, memErr         = mem.VirtualMemory()
		mMemUsed             float64
		mDisk, diskErr       = disk.Usage("/")
		mProcess, processErr = process.Pids()
		mLoadAvg             = new(LoadAvg)
		monitorHeads         []MonitorHead
		nets                 []NetC
		meta                 = getMonitor()
	)

	if cpuErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": cpuErr,
		}).Error("monitor")

		mCpu = []cpu.InfoStat{{VendorID: "", ModelName: ""}}
	}

	if memErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": memErr,
		}).Error("monitor")
		mMem = new(mem.VirtualMemoryStat)
	}

	if diskErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": diskErr,
		}).Error("monitor")
		mDisk = new(disk.UsageStat)
	}

	if processErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": processErr,
		}).Error("monitor")
	}

	// cpu使用率
	cu, err := cpu.Percent(time.Second, false)
	if err == nil {
		mCpuUsed = cast.ToFloat64(fmt.Sprintf("%.2f", cu[0]))
	}

	// 内存使用率
	mMemUsed = cast.ToFloat64(fmt.Sprintf("%.2f", mMem.UsedPercent))
	// 负载
	if len(meta.LoadAvg) > 0 {
		mLoadAvg = meta.LoadAvg[len(meta.LoadAvg)-1]
	}

	monitorHeads = append(monitorHeads, MonitorHead{
		Title:       "CPU",
		Data:        "使用率 " + Round2String(mCpuUsed) + "%",
		BottomTitle: "CPU数量",
		TotalSum:    cast.ToString(runtime.NumCPU()) + "核心",
		IconClass:   "Cpu",
		Extra: MonitorHeadExtra{
			Data:  mCpu[0].VendorID,
			Data1: mCpu[0].ModelName,
		}},
		MonitorHead{
			Title:       "内存",
			Data:        "使用率 " + Round2String(mMemUsed) + "%",
			BottomTitle: "总内存",
			TotalSum:    FileSize(int64(mMem.Total)),
			IconClass:   "PieChart",
			Extra: MonitorHeadExtra{
				Data:  FileSize(int64(mMem.Used)),
				Data1: FileSize(int64(mMem.Total - mMem.Used)),
			}},
		MonitorHead{
			Title:       "磁盘",
			Data:        "已用 " + FileSize(int64(mDisk.Used)),
			BottomTitle: "总容量",
			TotalSum:    FileSize(int64(mDisk.Total)),
			IconClass:   "PieChart",
			Extra: MonitorHeadExtra{
				Data: Round2String(mDisk.UsedPercent),
			}},
		MonitorHead{
			Title:       "负载",
			Data:        Round2String(mLoadAvg.Ratio) + "%",
			BottomTitle: "总进程数",
			TotalSum:    cast.ToString(len(mProcess)) + "个",
			IconClass:   "PieChart",
		})

	for _, v := range meta.NetIO {
		nets = append(nets, NetC{
			Time:      v.Time,
			BytesSent: FileSize(int64(v.BytesSent)), // 转换为最大整数单位
			BytesRecv: FileSize(int64(v.BytesRecv)),
			Down:      Round2Float64(v.Down / 1024), // 转换为kb
			UP:        Round2Float64(v.UP / 1024),
		})
	}
	data := map[string]interface{}{
		"head": monitorHeads,
		"load": meta.LoadAvg,
		"net":  nets,
	}

	newCtx.JSON(consts.StatusOK, utils.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": data,
	})
}
