package routes

import (
	"cloud/api/pay/app"
	"cloud/api/pay/channel"
	"cloud/api/pay/order"
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
	}
}
