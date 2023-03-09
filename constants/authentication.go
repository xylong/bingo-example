package constants

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	// ATokenExpired access_token过期时间
	ATokenExpired = time.Hour

	// ATokenExpireKey access_token配置名
	ATokenExpireKey = "jwt.accessExpire"

	// RTokenExpired refresh_token过期时间
	RTokenExpired = time.Hour * 24 * 3

	// RTokenExpiredKey refresh_token配置名
	RTokenExpiredKey = "jwt.refreshExpire"

	// SessionID 用户会话标识
	SessionID = "user_id"
)

// UserClaims 定义payload
type UserClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
