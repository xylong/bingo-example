package middleware

import (
	"bingo-example/application/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"net/http"
	"strings"
)

// Authentication 鉴权
type Authentication struct {
}

func (a *Authentication) Before(ctx *bingo.Context) error {
	token := ctx.Token()
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		return fmt.Errorf("unauthorized")
	}

	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})

		return fmt.Errorf("unauthorized")
	}

	claims, err := new(service.JwtService).ParseToken(parts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "access denied",
			"data":    nil,
		})
		return fmt.Errorf("unauthorized")
	}

	ctx.Set("user_id", claims.ID)
	ctx.Next()
	return nil
}
