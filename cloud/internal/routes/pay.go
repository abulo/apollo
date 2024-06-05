package routes

import (
	"cloud/api/pay/app"
	"cloud/api/pay/channel"
	"cloud/api/pay/notify"
	"cloud/api/pay/order"
	"cloud/api/pay/refund"
	"cloud/internal/middleware"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func PayInitRoute(handle *xhertz.Server) {
	auth := handle.Group("/admin").Use(middleware.AuthMiddleware())
	{
		// pay_app->支付应用信息->创建
		auth.POST("/pay/app", app.PayAppCreate)
		// pay_app->支付应用信息->更新
		auth.PUT("/pay/app/:id/update", app.PayAppUpdate)
		// pay_app->支付应用信息->删除
		auth.DELETE("/pay/app/:id/delete", app.PayAppDelete)
		// pay_app->支付应用信息->单条数据信息查看
		auth.GET("/pay/app/:id/item", app.PayApp)
		// pay_app->支付应用信息->恢复
		auth.PUT("/pay/app/:id/recover", app.PayAppRecover)
		// pay_app->支付应用信息->列表
		auth.GET("/pay/app", app.PayAppList)

		// pay_channel->支付渠道->创建
		auth.POST("/pay/app/:payAppId/channel/:channelCode", channel.PayChannelCreate)
		// pay_channel->支付渠道->单条数据信息查看
		auth.GET("/pay/app/:payAppId/channel/:channelCode/item", channel.PayChannel)

		// pay_order->支付订单->创建
		// auth.POST("/pay/order", order.PayOrderCreate)
		// pay_order->支付订单->更新
		// auth.PUT("/pay/order/:id/update", order.PayOrderUpdate)
		// pay_order->支付订单->删除
		auth.DELETE("/pay/order/:id/delete", order.PayOrderDelete)
		// pay_order->支付订单->单条数据信息查看
		auth.GET("/pay/order/:id/item", order.PayOrder)
		// pay_order->支付订单->恢复
		auth.PUT("/pay/order/:id/recover", order.PayOrderRecover)
		// pay_order->支付订单->列表
		auth.GET("/pay/order", order.PayOrderList)

		// pay_refund->退款订单->创建
		// auth.POST("/pay/refund", refund.PayRefundCreate)
		// pay_refund->退款订单->更新
		// auth.PUT("/pay/refund/:id/update", refund.PayRefundUpdate)
		// pay_refund->退款订单->删除
		auth.DELETE("/pay/refund/:id/delete", refund.PayRefundDelete)
		// pay_refund->退款订单->单条数据信息查看
		auth.GET("/pay/refund/:id/item", refund.PayRefund)
		// pay_refund->退款订单->恢复
		auth.PUT("/pay/refund/:id/recover", refund.PayRefundRecover)
		// pay_refund->退款订单->列表
		auth.GET("/pay/refund", refund.PayRefundList)

		// pay_notify_task->商户支付-任务通知->创建
		// auth.POST("/pay/notify", notify.PayNotifyTaskCreate)
		// pay_notify_task->商户支付-任务通知->更新
		// auth.PUT("/pay/notify/:id/update", notify.PayNotifyTaskUpdate)
		// pay_notify_task->商户支付-任务通知->删除
		auth.DELETE("/pay/notify/:id/delete", notify.PayNotifyTaskDelete)
		// pay_notify_task->商户支付-任务通知->单条数据信息查看
		auth.GET("/pay/notify/:id/item", notify.PayNotifyTask)
		// pay_notify_task->商户支付-任务通知->恢复
		auth.PUT("/pay/notify/:id/recover", notify.PayNotifyTaskRecover)
		// pay_notify_task->商户支付-任务通知->列表
		auth.GET("/pay/notify", notify.PayNotifyTaskList)
	}
}
