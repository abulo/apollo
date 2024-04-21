package dao

import (
	"github.com/golang-jwt/jwt/v5"
)

type SystemUserLogin struct {
	Username    string `json:"username,required"`    // 用户名
	Password    string `json:"password,required"`    // 密码
	CaptchaCode string `json:"captchaCode,required"` // 验证码
	CaptchaId   string `json:"captchaId,required"`   // 验证码ID
}

// SystemUserToken 用户令牌
type SystemUserToken struct {
	UserId   int64  `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	NickName string `json:"nickName"` // 用户昵称
	TenantId int64  `json:"tenantId"` // 租户ID
	jwt.RegisteredClaims
}

// 用户密码
type SystemUserPassword struct {
	Password string `json:"password,required"` // 密码
}
