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
		ExpiresAt: getJWTTime(constants.ATokenExpired),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	ac := &constants.UserClaims{
		ID:               id,
		RegisteredClaims: rc,
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, ac).SignedString(secret); err != nil {
		return
	}

	rc.ExpiresAt = getJWTTime(constants.RTokenExpired)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(secret)

	return
}

// Parse 解析token
func Parse(token string) (*constants.UserClaims, error) {
	claims := &constants.UserClaims{}
	_token, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if err != nil || !_token.Valid {
		err = fmt.Errorf("invalid token")
	}

	return claims, err
}
