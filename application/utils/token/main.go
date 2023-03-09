package token

import (
	"bingo-example/constants"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Generate 生成token
// access_token 用于访问token
// refresh_token 用于刷新access_token
func Generate(id int) (accessToken, refreshToken string, err error) {
	secret := getSecret()

	rc := jwt.RegisteredClaims{
		ExpiresAt: getExpiryDate(constants.ATokenExpireKey, getJWTTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	ac := &constants.UserClaims{
		ID:               id,
		RegisteredClaims: rc,
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, ac).SignedString(secret); err != nil {
		return
	}

	rc.ExpiresAt = getExpiryDate(constants.RTokenExpiredKey, getJWTTime)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(secret)

	return
}

// Parse 解析token
func Parse(token string) (*constants.UserClaims, error) {
	claims := &constants.UserClaims{}
	_token, err := jwt.ParseWithClaims(token, claims, key)

	if err != nil {
		return nil, err
	}

	if !_token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// Refresh 刷新token
func Refresh(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// refreshToken无效直接返回
	if _, err = jwt.Parse(refreshToken, key); err != nil {
		return
	}

	// 从旧accessToken解析出claims
	claims := &constants.UserClaims{}
	_, err = jwt.ParseWithClaims(accessToken, claims, key)

	// 如果是正常过期则重新生成
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return Generate(claims.ID)
	}

	return
}
