package dao

import (
	"github.com/golang-jwt/jwt/v5"
)

type SystemUserLogin struct {
	Username   string `json:"username,required"`   // 用户名
	Password   string `json:"password,required"`   // 密码
	VerifyCode string `json:"verifyCode,required"` // 验证码
	VerifyId   string `json:"verifyId,required"`   // 验证码ID
}

// SystemUserToken 用户令牌
type SystemUserToken struct {
	SystemUserId   int64  `json:"systemUserId"`   // 用户ID
	SystemUserName string `json:"systemUserName"` // 用户名
	SystemNickName string `json:"systemNickName"` // 用户昵称
	jwt.RegisteredClaims
}

// 用户密码
type SystemUserPassword struct {
	Password string `json:"password,required"` // 密码
}
