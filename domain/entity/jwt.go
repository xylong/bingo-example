package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// JwtSecret 密钥
var JwtSecret = []byte("bingo")

// JwtClaims jwt数据
type JwtClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
