package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// 获取过期时间
func getJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

// 获取jwt密钥
func getSecret() []byte {
	return []byte(viper.GetString("jwt.secret"))
}
