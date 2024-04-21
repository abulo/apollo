package logger

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/logger"
	"context"
	"encoding/json"
	"time"

	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SystemEntryLogDelete 删除数据
func SystemEntryLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	var reqInfo dao.SystemEntryLogDelete
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	var mongoIds []string
	if err := json.Unmarshal(reqInfo.Ids.JSON, &mongoIds); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	systemEntryLogIds := make([]primitive.ObjectID, 0)
	for _, mongoId := range mongoIds {
		id, _ := primitive.ObjectIDFromHex(mongoId)
		systemEntryLogIds = append(systemEntryLogIds, id)
	}
	if len(systemEntryLogIds) == 0 {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	request := make(bson.D, 0)
	request = util.ConvertBson(request, bson.E{Key: "_id", Value: bson.D{{Key: "$in", Value: systemEntryLogIds}}})
	if _, err := logger.SystemEntryLogDelete(ctx, request); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
	})
}

// SystemEntryLogDrop 删除数据
func SystemEntryLogDrop(ctx context.Context, newCtx *app.RequestContext) {
	var reqInfo dao.SystemEntryLogDrop
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request := make(bson.D, 0)
	if reqInfo.StartTime.IsValid() {
		request = util.ConvertBson(request, bson.E{Key: "timestamp", Value: bson.D{{Key: "$gte", Value: *reqInfo.StartTime.Ptr()}}})
	}
	if reqInfo.EndTime.IsValid() {
		request = util.ConvertBson(request, bson.E{Key: "timestamp", Value: bson.D{{Key: "$lte", Value: *reqInfo.EndTime.Ptr()}}})
	}
	if _, err := logger.SystemEntryLogDrop(ctx, request); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
	})
}

// SystemEntryLogList 获取数据
func SystemEntryLogList(ctx context.Context, newCtx *app.RequestContext) {
	request := make(bson.D, 0)
	if val, ok := newCtx.GetQuery("startTime"); ok {
		request = util.ConvertBson(request, bson.E{
			Key: "timestamp",
			Value: bson.D{
				{
					Key:   "$gte",
					Value: cast.ToTimeInDefaultLocation(val, time.Local),
				},
			},
		})
	}
	if val, ok := newCtx.GetQuery("endTime"); ok {
		request = util.ConvertBson(request, bson.E{
			Key: "timestamp",
			Value: bson.D{
				{
					Key:   "$lte",
					Value: cast.ToTimeInDefaultLocation(val, time.Local),
				},
			},
		})
	}
	pageNum := cast.ToInt64(newCtx.Query("pageNum"))
	pageSize := cast.ToInt64(newCtx.Query("pageSize"))
	total, _ := logger.SystemEntryLogListTotal(ctx, request)
	list, _ := logger.SystemEntryLogList(ctx, request, pageNum, pageSize)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": 200,
		"msg":  "OK",
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  pageNum,
			"pageSize": pageSize,
		},
	})
}

// SystemEntryLog 删除数据
func SystemEntryLog(ctx context.Context, newCtx *app.RequestContext) {
	systemEntryId := cast.ToString(newCtx.Param("id"))
	// 将systemEntryId 装换成 primitive.ObjectID 类型
	id, _ := primitive.ObjectIDFromHex(systemEntryId)
	res, err := logger.SystemEntryLog(ctx, id)
	if err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": res,
	})
}
