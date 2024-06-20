package pay

import (
	"cloud/service/pay/app"
	"cloud/service/pay/channel"
	"cloud/service/pay/notify"
	"cloud/service/pay/order"
	"cloud/service/pay/refund"
	"cloud/service/pay/wallet"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 支付应用信息->pay_app
	app.RegisterPayAppServiceServer(server.Server, &app.SrvPayAppServiceServer{
		Server: server,
	})
	// 支付渠道->pay_channel
	channel.RegisterPayChannelServiceServer(server.Server, &channel.SrvPayChannelServiceServer{
		Server: server,
	})
	// 支付订单->pay_order
	order.RegisterPayOrderServiceServer(server.Server, &order.SrvPayOrderServiceServer{
		Server: server,
	})
	// 支付订单拓展->pay_order_extension
	order.RegisterPayOrderExtensionServiceServer(server.Server, &order.SrvPayOrderExtensionServiceServer{
		Server: server,
	})
	// 商户支付-任务通知->pay_notify_task
	notify.RegisterPayNotifyTaskServiceServer(server.Server, &notify.SrvPayNotifyTaskServiceServer{
		Server: server,
	})
	// 支付通知日志->pay_notify_log
	notify.RegisterPayNotifyLogServiceServer(server.Server, &notify.SrvPayNotifyLogServiceServer{
		Server: server,
	})
	// 退款订单->pay_refund
	refund.RegisterPayRefundServiceServer(server.Server, &refund.SrvPayRefundServiceServer{
		Server: server,
	})
	// 会员钱包表->pay_wallet
	wallet.RegisterPayWalletServiceServer(server.Server, &wallet.SrvPayWalletServiceServer{
		Server: server,
	})
	// 充值套餐表->pay_wallet_recharge_package
	wallet.RegisterPayWalletRechargePackageServiceServer(server.Server, &wallet.SrvPayWalletRechargePackageServiceServer{
		Server: server,
	})
}
