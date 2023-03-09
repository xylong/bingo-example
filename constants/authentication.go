package constants

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	// ATokenExpired access_token过期时间
	ATokenExpired = time.Hour

	// RTokenExpired refresh_token过期时间
	RTokenExpired = time.Hour * 24 * 3

	// EffectTime 默认token有效期
	EffectTime = 60

	// SessionID 用户会话标识
	SessionID = "user_id"
)

// UserClaims 定义payload
type UserClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
