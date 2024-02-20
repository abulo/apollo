package tools

import (
	"cloud/dao"
	"errors"
	"time"

	"github.com/abulo/ratel/v3/util"
	"github.com/golang-jwt/jwt/v5"
)

// 签名密钥
const signKey = "hezhian"

// GenerateToken 生成token
func GenerateToken(id int64, userName, nickName, Audience string) (string, error) {
	claim := dao.SystemUserToken{
		SystemUserId:   id,
		SystemUserName: userName,
		SystemNickName: nickName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AuthServer",                                      // 签发者
			Subject:   "Auth",                                            // 签发对象
			Audience:  jwt.ClaimStrings{Audience},                        // 签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 签发时间
			ID:        util.Random(),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(signKey))
	return token, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*dao.SystemUserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dao.SystemUserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}
	claims, ok := token.Claims.(*dao.SystemUserToken)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
