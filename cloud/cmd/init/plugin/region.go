package plugin

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/region"
	"context"
	"encoding/json"
	"regexp"

	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

var (
	historyMap map[string]string
	weatherMap map[string]any
	regionList []*dao.Region
)

func InitRegionJson() error {
	historyUrl := initial.Core.Path + "/config/json/region.json"
	jsonContent, err := util.FileGetContents(historyUrl)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonContent), &historyMap)
}

func InitWeatherJson() error {
	weatherUrl := initial.Core.Path + "/config/json/weather.json"
	jsonContent, err := util.FileGetContents(weatherUrl)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonContent), &weatherMap)
}

// 初始化区域
func InitRegion() {
	if err := InitRegionJson(); err != nil {
		logger.Logger.Info("InitRegionJson error: ", err)
		return
	}
	if err := InitWeatherJson(); err != nil {
		logger.Logger.Info("InitWeatherJson error: ", err)
		return
	}

	regionList := InitProvince()
	// for _, regionInfo := range regionList {
	// 	// 判断直辖市和想过澳门地区
	// 	if *regionInfo.Id == 110000 || *regionInfo.Id == 120000 || *regionInfo.Id == 310000 || *regionInfo.Id == 500000 || *regionInfo.Id == 810000 || *regionInfo.Id == 820000 {
	// 		ZhiXiaShiWeather(regionInfo)
	// 	} else {

	// 		ProvinceWeather(regionInfo)
	// 	}
	// }
	ctx := context.Background()
	// 插入数据库
	InsertRegion(ctx, regionList)
}

func InsertRegion(ctx context.Context, regionList []*dao.Region) {
	for _, regionInfo := range regionList {
		_, err := region.RegionCreate(ctx, *regionInfo)
		if err != nil {
			logger.Logger.Info("RegionCreate error: ", err)
		}
		if len(regionInfo.Children) > 0 {
			InsertRegion(ctx, regionInfo.Children)
		}
	}
}

func InitProvince() []*dao.Region {
	// 先遍历一遍，将所有省级地区放入 regionList
	pattern := regexp.MustCompile(`0000$`)
	for code, val := range historyMap {
		if pattern.MatchString(code) {
			// 获取市级地区
			regionList = append(regionList, &dao.Region{
				Id:       proto.Int64(cast.ToInt64(code)),
				Name:     proto.String(getName(val)),
				ParentId: proto.Int64(0),
			})
		} else {
			// 放入地市
			// 首先用正则替换最后两位为00
			codeCity := regexp.MustCompile(`\d{2}$`).ReplaceAllString(code, "00")
			// fmt.Println(codeCity)
			codeDistrict := regexp.MustCompile(`\d{4}$`).ReplaceAllString(code, "0000")
			if _, cityOk := historyMap[codeCity]; cityOk {
				// 获取市级地区
				if codeCity != code {
					regionList = append(regionList, &dao.Region{
						Id:       proto.Int64(cast.ToInt64(code)),
						Name:     proto.String(getName(val)),
						ParentId: proto.Int64(cast.ToInt64(codeCity)),
					})
				} else {
					regionList = append(regionList, &dao.Region{
						Id:       proto.Int64(cast.ToInt64(code)),
						Name:     proto.String(getName(val)),
						ParentId: proto.Int64(cast.ToInt64(codeDistrict)),
					})
				}
			} else {
				if _, districtOk := historyMap[codeDistrict]; districtOk {
					regionList = append(regionList, &dao.Region{
						Id:       proto.Int64(cast.ToInt64(code)),
						Name:     proto.String(getName(val)),
						ParentId: proto.Int64(cast.ToInt64(codeDistrict)),
					})
				}
			}
		}
	}
	return RegionTree(regionList, 0)
}

func RegionTree(list []*dao.Region, parentId int64) []*dao.Region {
	var tree []*dao.Region
	for _, item := range list {
		if *item.ParentId == parentId {
			item.Children = RegionTree(list, *item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}

func getName(name string) string {
	return name
}

// // 处理省市区的天气
// func ProvinceWeather(regionInfo *dao.Region) {
// 	// 遍历省级地区
// 	for provinceName, weatherItem := range weatherMap {
// 		// 判断字符串中知否包含指定的字符串
// 		if strings.Contains(*regionInfo.Name, provinceName) {
// 			cityMap := cast.ToStringMap(weatherItem)
// 			cityChildren := regionInfo.Children
// 			for cityName, cityItem := range cityMap {
// 				// 遍历children信息
// 				for _, districtItem := range cityChildren {
// 					if strings.Contains(*districtItem.Name, cityName) {
// 						cityWeather := cast.ToStringMap(cityItem)
// 						cityWeatherInfo := cast.ToStringMap(cityWeather[cityName])
// 						districtItem.WeatherCode = null.StringFrom(cast.ToString(cityWeatherInfo["AREAID"]))
// 						// 县级
// 						xianChildren := districtItem.Children
// 						for xianName, xianItem := range cityWeather {
// 							for _, xianVal := range xianChildren {
// 								if strings.Contains(*xianVal.Name, xianName) {
// 									xianWeather := cast.ToStringMap(xianItem)
// 									xianVal.WeatherCode = null.StringFrom(cast.ToString(xianWeather["AREAID"]))
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// // ZhiXiaShiWeather 处理直辖市的天气
// func ZhiXiaShiWeather(regionInfo *dao.Region) {
// 	for provinceName, weatherItem := range weatherMap {
// 		// 判断字符串中知否包含指定的字符串
// 		if strings.Contains(*regionInfo.Name, provinceName) {
// 			// 因为直辖市只有一个层级
// 			cityList := cast.ToStringMap(cast.ToStringMap(weatherItem)[provinceName])
// 			// 需要将第一个层级的WeatherCode替换掉
// 			childrenList := regionInfo.Children
// 			for cityName, cityItem := range cityList {
// 				if provinceName == cityName {
// 					// 第一个层级, 北京,天津
// 					provinceWeather := cast.ToStringMap(cityItem)
// 					regionInfo.WeatherCode = null.StringFrom(cast.ToString(provinceWeather["AREAID"]))
// 				}
// 				// 遍历children信息
// 				for _, districtItem := range childrenList {
// 					if strings.Contains(*districtItem.Name, cityName) {
// 						// districtItem.Name = proto.String(cityName)
// 						provinceWeather := cast.ToStringMap(cityItem)
// 						districtItem.WeatherCode = null.StringFrom(cast.ToString(provinceWeather["AREAID"]))
// 					}
// 				}
// 			}
// 		}
// 	}

// }
