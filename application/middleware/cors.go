package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 跨域
type Cors struct {
}

func NewCors() *Cors {
	return &Cors{}
}

func (c *Cors) Before(ctx *gin.Context) error {
	method := ctx.Request.Method

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	ctx.Next()
	return nil
}

func (c *Cors) After(data interface{}) (interface{}, error) {
	return data, nil
}
