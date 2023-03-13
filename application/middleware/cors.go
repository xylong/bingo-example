package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域
type cors struct {
}

func newCors() *cors {
	return &cors{}
}

func (c *cors) Name() string {
	return "cors"
}

func (c *cors) Before(ctx *gin.Context) error {
	fmt.Println("before cors")
	method := ctx.Request.Method

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	return nil
}

func (c *cors) After(data interface{}) (interface{}, error) {
	fmt.Println("after cors")
	return data, nil
}
