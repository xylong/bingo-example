package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 请求参数校验
type request struct {
}

func newRequest() *request {
	return &request{}
}

func (r *request) Name() string {
	return "request"
}

func (r *request) Before(ctx *gin.Context) error {
	fmt.Println("before req")
	return nil
}

func (r *request) After(data interface{}) (interface{}, error) {
	fmt.Println("after req")
	return data, nil
}
