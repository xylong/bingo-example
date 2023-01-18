package middleware

import (
	"bingo-example/application/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xylong/bingo"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// Authentication 鉴权
type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (a *Authentication) Before(ctx *bingo.Context) error {
	var err error

	defer func() {
		if err != nil {
			zap.L().Warn(err.Error(),
				zap.String("path", ctx.Request.URL.Path),
				zap.String("method", ctx.Request.Method),
				zap.String("ip", ctx.ClientIP()))
		}
	}()

	token := ctx.Token()
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		err = errors.New("unauthorized")
		return nil
	}

	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		err = jwt.ErrTokenMalformed
		return nil
	}

	claims, err := new(service.JwtService).ParseToken(parts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		return nil
	}

	ctx.Set("user_id", claims.ID)
	ctx.Next()
	return nil
}

func (a *Authentication) After(data interface{}) (interface{}, error) {
	return data, nil
}
