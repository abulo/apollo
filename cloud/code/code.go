package code

import "google.golang.org/grpc/codes"

// 1000～1999 区间表示系统错误
// 2000～2999 区间表示用户错误
var (
	Success               = int64(200)  // 执行成功
	SystemError           = int64(1001) // 系统错误
	ResourceNotAvailable  = int64(1002) // 服务端资源不可用
	RemoteServiceError    = int64(1003) // 远程服务出错
	ParamInvalid          = int64(1004) // 错误:参数错误
	SystemBusy            = int64(1005) // 任务过多系统繁忙
	Timeout               = int64(1006) // 任务超时
	RPCError              = int64(1007) // RPC错误
	BadRequest            = int64(1008) // 非法请求
	VerifyCodeError       = int64(1009) // 验证码错误
	TokenError            = int64(1010) // Token错误
	TokenEmptyError       = int64(1011) // Token为空
	TokenInvalidError     = int64(1012) // Token无效
	TokenExpiredError     = int64(1013) // Token 已经过期
	TokenNotValidYetError = int64(1014) // Token 未激活
	TokenMalformedError   = int64(1015) // Token 错误
	NotFound              = int64(1016) // 未找到
	LoginFail             = int64(1017) // 登录失败
	SqlError              = int64(2001) // SQL错误
	RedisError            = int64(2002) // Redis错误
	UnavailableError      = int64(2003) // 数据不可用
	NoPermission          = int64(2004) // 无权限
	EncryptError          = int64(2005) // 加密错误
	DecryptError          = int64(2006) // 解密错误
	DeptError             = int64(2007) // 部门未创建

	//状态码对应的信息
	statusText = map[int64]string{
		Success:               "成功",
		SystemError:           "系统错误",
		ResourceNotAvailable:  "服务端资源不可用",
		RemoteServiceError:    "远程服务出错",
		ParamInvalid:          "参数错误",
		SystemBusy:            "任务过多系统繁忙",
		Timeout:               "任务超时",
		RPCError:              "RPC错误",
		BadRequest:            "非法请求",
		SqlError:              "SQL错误",
		RedisError:            "Redis错误",
		VerifyCodeError:       "验证码错误",
		TokenError:            "Token错误",
		TokenEmptyError:       "Token为空",
		TokenInvalidError:     "Token无效",
		TokenExpiredError:     "Token已经过期",
		TokenNotValidYetError: "Token未激活",
		TokenMalformedError:   "Token错误",
		NotFound:              "未找到",
		LoginFail:             "登录失败",
		UnavailableError:      "数据不可用",
		NoPermission:          "无权限",
		EncryptError:          "加密错误",
		DecryptError:          "解密错误",
		DeptError:             "部门未创建",
	}
)

func StatusText(code int64) string {
	if val, ok := statusText[code]; ok {
		return val
	}
	return "未知错误"
}

func ConvertToGrpc(code int64) codes.Code {
	return codes.Code(code)
}

func ConvertToHttp(code codes.Code) int64 {
	return int64(code)
}
