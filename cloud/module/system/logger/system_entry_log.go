package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SystemEntryLogList 获取系统日志列表
func SystemEntryLogList(ctx context.Context, condition bson.D, pageNum, pageSize int64) (res []dao.SystemEntryLog, err error) {
	skip := pageSize * (pageNum - 1)
	sort := make(bson.D, 0)
	sort = util.ConvertBson(sort, bson.E{Key: "_id", Value: -1})
	model, err := initial.Core.Store.LoadMongoDB("mongodb").NewModel("logger_entry")
	if err != nil {
		return
	}
	list := make([]map[string]interface{}, 0)
	err = model.Collection.Where(condition).Sort(sort).Skip(skip).Limit(pageSize).FindMany(ctx, &list)
	if err == nil {
		for _, v := range list {
			res = append(res, dao.SystemEntryLog{
				Id:        cast.ToString(v["_id"].(primitive.ObjectID).Hex()),
				Host:      cast.ToString(v["host"]),
				Timestamp: util.Date("Y-m-d H:i:s", v["timestamp"].(primitive.DateTime).Time()),
				File:      cast.ToString(v["file"]),
				Func:      cast.ToString(v["func"]),
				Message:   cast.ToString(v["message"]),
				Data:      v["data"],
				Level:     cast.ToString(v["level"]),
			})
		}
	}
	return
}

// SystemEntryLogListTotal 获取系统日志列表总数
func SystemEntryLogListTotal(ctx context.Context, condition bson.D) (res int64, err error) {
	model, err := initial.Core.Store.LoadMongoDB("mongodb").NewModel("logger_entry")
	if err != nil {
		return
	}
	res, err = model.Collection.Where(condition).Count(ctx)
	return
}

// SystemEntryLogDelete 删除系统日志
func SystemEntryLogDelete(ctx context.Context, condition bson.D) (res int64, err error) {
	model, err := initial.Core.Store.LoadMongoDB("mongodb").NewModel("logger_entry")
	if err != nil {
		return
	}
	res, err = model.Collection.Where(condition).Delete(ctx)
	return
}

// SystemEntryLogDrop 清空系统日志
func SystemEntryLogDrop(ctx context.Context, condition bson.D) (res int64, err error) {
	res = 0
	model, err := initial.Core.Store.LoadMongoDB("mongodb").NewModel("logger_entry")
	if err != nil {
		return
	}
	if len(condition) == 0 || condition == nil {
		err = model.Collection.Drop(ctx)
	} else {
		_, err = model.Collection.Where(condition).Delete(ctx)
	}
	return
}

func SystemEntryLog(ctx context.Context, systemEntryId primitive.ObjectID) (res dao.SystemEntryLog, err error) {
	model, err := initial.Core.Store.LoadMongoDB("mongodb").NewModel("logger_entry")
	if err != nil {
		return
	}
	condition := make(bson.D, 0)
	condition = util.ConvertBson(condition, bson.E{Key: "_id", Value: systemEntryId})
	var list map[string]interface{}
	err = model.Collection.Where(condition).FindOne(ctx, &list)
	if err == nil {
		res = dao.SystemEntryLog{
			Id:        cast.ToString(list["_id"].(primitive.ObjectID).Hex()),
			Host:      cast.ToString(list["host"]),
			Timestamp: util.Date("Y-m-d H:i:s", list["timestamp"].(primitive.DateTime).Time()),
			File:      cast.ToString(list["file"]),
			Func:      cast.ToString(list["func"]),
			Message:   cast.ToString(list["message"]),
			Data:      list["data"],
			Level:     cast.ToString(list["level"]),
		}
	}
	return
}
