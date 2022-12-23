package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo"
	"go.uber.org/zap"
	"time"
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
	_ = bingo.RegisterBindTag("date", CheckDate)
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

// CheckDate 检测日期格式
func CheckDate(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, err := time.Parse("2006-01-02", v)
	return err == nil
}
