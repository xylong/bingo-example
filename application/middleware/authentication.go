package middleware

import (
	"bingo-example/application/utils/token"
	"bingo-example/constants"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// authentication 鉴权
type authentication struct{}

func newAuthentication() *authentication {
	return &authentication{}
}

func (a *authentication) Name() string {
	return "auth"
}

func (a *authentication) Before(ctx *gin.Context) error {
	fmt.Println("before auth")
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

func (a *authentication) After(data interface{}) (interface{}, error) {
	fmt.Println("after auth")
	return data, nil
}
