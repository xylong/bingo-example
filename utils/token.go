package utils

import (
	"bingo-example/constants"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// GenerateToken 生成token
func GenerateToken(id int) (string, error) {
	expiration := getExpire()
	now := time.Now()

	claims := &constants.UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expiration * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 用指定方法创建签名对象
	return token.SignedString(getSecret())                     // 使用指定的secret签名并获得完整的编码后的字符串token
}

// ParseToken 解析token
func ParseToken(token string) (*constants.UserClaims, error) {
	claims := &constants.UserClaims{}
	_token, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if err != nil || !_token.Valid {
		err = fmt.Errorf("invalid token")
	}

	return claims, err
}

// 获取jwt密钥
func getSecret() []byte {
	return []byte(viper.GetString("jwt.secret"))
}

// 获取token有效期
func getExpire() time.Duration {
	expiration := viper.GetDuration("jwt.tokenExpire")
	if expiration == 0 {
		expiration = constants.EffectTime
	}

	return expiration
}
