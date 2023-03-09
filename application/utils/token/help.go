package token

import (
	"bingo-example/constants"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type numericDate func(time.Duration) *jwt.NumericDate

// 获取过期时间
func getJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

// getExpiryDate 获取token有效期
func getExpiryDate(key string, date numericDate) *jwt.NumericDate {
	n := viper.GetDuration(key) * time.Minute

	if n <= 0 {
		if key == constants.ATokenExpireKey {
			n = constants.ATokenExpired
		} else {
			n = constants.RTokenExpired
		}
	}

	return date(n)
}

// 获取jwt密钥
func getSecret() []byte {
	return []byte(viper.GetString("jwt.secret"))
}
