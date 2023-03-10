package middleware

import (
	"bingo-example/application/utils/token"
	"bingo-example/constants"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// Authentication 鉴权
type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (a *Authentication) Before(ctx *gin.Context) error {
	var err error

	defer func() {
		if err != nil {
			zap.L().Warn(err.Error(),
				zap.String("path", ctx.Request.URL.Path),
				zap.String("method", ctx.Request.Method),
				zap.String("ip", ctx.ClientIP()))
		}
	}()

	tokenStr := ctx.Request.Header.Get("Authorization")
	if tokenStr == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		err = errors.New("unauthorized")
		return nil
	}

	parts := strings.SplitN(tokenStr, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		err = jwt.ErrTokenMalformed
		return nil
	}

	claims, err := token.Parse(parts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		return nil
	}

	ctx.Set(constants.SessionID, claims.ID)
	ctx.Next()
	return nil
}

func (a *Authentication) After(data interface{}) (interface{}, error) {
	return data, nil
}
