package middlewares

import (
	"bingo-example/pkg/helpers"
	"bingo-example/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io"
	"time"
)

type log struct {
}

func newLog() *log {
	return &log{}
}

func (l *log) Before(ctx *gin.Context) error {
	// 获取 response 内容
	w := &responseBodyWriter{
		body:           &bytes.Buffer{},
		ResponseWriter: ctx.Writer,
	}
	ctx.Writer = w

	// 获取请求数据
	var requestBody []byte
	if ctx.Request.Body != nil {
		// c.Request.Body 是一个 buffer 对象，只能读取一次
		requestBody, _ = io.ReadAll(ctx.Request.Body)
		// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	// 设置开始时间
	start := time.Now()
	ctx.Next()

	// 开始记录日志的逻辑
	cost := time.Since(start)
	responseStatus := ctx.Writer.Status()

	logFields := []zap.Field{
		zap.Int("status", responseStatus),
		zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
		zap.String("query", ctx.Request.URL.RawQuery),
		zap.String("ip", ctx.ClientIP()),
		zap.String("user-agent", ctx.Request.UserAgent()),
		zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.String("time", helpers.MicrosecondsStr(cost)),
	}

	if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
		// 请求的内容
		logFields = append(logFields, zap.String("Request Body", string(requestBody)))

		// 响应的内容
		logFields = append(logFields, zap.String("Response Body", w.body.String()))
	}

	if responseStatus > 400 && responseStatus <= 499 {
		// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
		logger.Warn("HTTP Warning "+cast.ToString(responseStatus), logFields...)
	} else if responseStatus >= 500 && responseStatus <= 599 {
		// 除了内部错误，记录 error
		logger.Error("HTTP Error "+cast.ToString(responseStatus), logFields...)
	} else {
		logger.Debug("HTTP Access Log", logFields...)
	}

	return nil
}

func (l *log) After(data interface{}) (interface{}, error) {
	return data, nil
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
