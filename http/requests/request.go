package requests

import (
	"bingo-example/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(*gin.Context, interface{}) map[string][]string

// Validate 参数验证
func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) gin.H {

	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := ctx.ShouldBind(obj); err != nil {
		return response.BadRequest(err)
	}

	// 2. 表单验证
	errs := handler(ctx, obj)

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		return response.ValidationError(ctx, errs)
	}

	return nil
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
