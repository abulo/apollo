package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// InitSql load mysql && returns an mysql instance.
func (initial *Initial) InitSql() *Initial {
	configs := initial.Config.Get("db")
	list := configs.(map[string]interface{})
	links := make(map[string]sql.SqlConn)
	for node, nodeConfig := range list {
		opts := make([]sql.Option, 0)
		res := nodeConfig.(map[string]interface{})
		if Host := cast.ToString(res["Host"]); Host != "" {
			opts = append(opts, sql.WithHost(Host))
		}
		if Port := cast.ToString(res["Port"]); Port != "" {
			opts = append(opts, sql.WithPort(Port))
		}
		if Username := cast.ToString(res["Username"]); Username != "" {
			opts = append(opts, sql.WithUsername(Username))
		}
		if Password := cast.ToString(res["Password"]); Password != "" {
			opts = append(opts, sql.WithPassword(Password))
		}
		if Charset := cast.ToString(res["Charset"]); Charset != "" {
			opts = append(opts, sql.WithCharset(Charset))
		}
		if Database := cast.ToString(res["Database"]); Database != "" {
			opts = append(opts, sql.WithDatabase(Database))
		}
		if MaxOpenConns := cast.ToInt(res["MaxOpenConns"]); MaxOpenConns > 0 {
			opts = append(opts, sql.WithMaxOpenConns(MaxOpenConns))
		}
		if MaxIdleConns := cast.ToInt(res["MaxIdleConns"]); MaxIdleConns > 0 {
			opts = append(opts, sql.WithMaxIdleConns(MaxIdleConns))
		}
		if MaxLifetime := cast.ToInt(res["MaxLifetime"]); MaxLifetime > 0 {
			opts = append(opts, sql.WithMaxLifetime(time.Duration(MaxLifetime)*time.Minute))
		}
		if MaxIdleTime := cast.ToInt(res["MaxIdleTime"]); MaxIdleTime > 0 {
			opts = append(opts, sql.WithMaxIdleTime(time.Duration(MaxIdleTime)*time.Minute))
		}
		opts = append(opts, sql.WithDisableMetric(cast.ToBool(res["DisableMetric"])))
		opts = append(opts, sql.WithDisableTrace(cast.ToBool(res["DisableTrace"])))
		if TimeZone := cast.ToString(res["TimeZone"]); TimeZone != "" {
			opts = append(opts, sql.WithTimeZone(TimeZone))
		}
		if DriverName := cast.ToString(res["DriverName"]); DriverName != "" {
			opts = append(opts, sql.WithDriverName(DriverName))
		}
		opts = append(opts, sql.WithDisablePrepare(cast.ToBool(res["DisablePrepare"])))
		opts = append(opts, sql.WithParseTime(cast.ToBool(res["ParseTime"])))
		if Addr := cast.ToStringSlice(res["Addr"]); len(Addr) > 0 {
			opts = append(opts, sql.WithAddr(Addr))
		}
		if DialTimeout := cast.ToString(res["DialTimeout"]); DialTimeout != "" {
			opts = append(opts, sql.WithDialTimeout(DialTimeout))
		}
		if OpenStrategy := cast.ToString(res["OpenStrategy"]); OpenStrategy != "" {
			opts = append(opts, sql.WithOpenStrategy(OpenStrategy))
		}
		opts = append(opts, sql.WithCompress(cast.ToBool(res["Compress"])))
		if MaxExecutionTime := cast.ToString(res["MaxExecutionTime"]); MaxExecutionTime != "" {
			opts = append(opts, sql.WithMaxExecutionTime(MaxExecutionTime))
		}
		opts = append(opts, sql.WithDisableDebug(cast.ToBool(res["DisableDebug"])))
		client, err := sql.NewClient(opts...)
		if err != nil {
			panic(err)
		}
		links["db."+node] = client.NewSqlClient()
	}

	proxyConfigs := initial.Config.Get("proxydb")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewSQL()
		if Master := cast.ToStringSlice(val["Master"]); len(Master) > 0 {
			for _, v := range Master {
				proxyPool.SetWrite(links[v])
			}
		}
		if Slave := cast.ToStringSlice(val["Slave"]); len(Slave) > 0 {
			for _, v := range Slave {
				proxyPool.SetRead(links[v])
			}
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Store.StoreSQL(Name, proxyPool)
		}
	}
	return initial
}
