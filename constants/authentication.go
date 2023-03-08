package constants

import "github.com/golang-jwt/jwt/v4"

const (
	// DefaultTokenExpire 默认token有效期
	DefaultTokenExpire = 60
)

// JwtClaims 定义payload
type JwtClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
