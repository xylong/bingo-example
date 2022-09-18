package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo"
	"go.uber.org/zap"
)

var valid *validator.Validate

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid = v
	} else {
		zap.L().Fatal("binding validator error")
	}

	_ = bingo.RegisterBindTag("phone", CheckPhone)
	_ = bingo.RegisterBindTag("sms", Field("required").toFunc())
}

var (
	// CheckPhone 检查手机号
	CheckPhone validator.Func = func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(string)
		return ok && len(v) == 11
	}

	// CheckSms 检查短信码
	CheckSms validator.Func = func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().([]string)
		return ok && len(v) == 6
	}
)

// Field 查询字段
type Field string

func (f Field) toFunc() validator.Func {
	bingo.ValidateMessage["field"] = "验证码必填"
	return CheckSms
}
