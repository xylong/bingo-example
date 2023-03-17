package auth

import (
	v1 "bingo-example/http/controllers/api/v1"
	"bingo-example/http/requests"
	"bingo-example/pkg/captcha"
	"bingo-example/pkg/logger"
	"bingo-example/pkg/response"
	"bingo-example/pkg/verifycode"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

func init() {
	v1.RegisterCtrl(NewVerifyCodeController())
}

// VerifyCodeController 验证码
type VerifyCodeController struct {
	v1.BaseApiController
}

func NewVerifyCodeController() *VerifyCodeController {
	return &VerifyCodeController{}
}

func (c *VerifyCodeController) Name() string {
	return "VerifyCodeController"
}

func (c *VerifyCodeController) showCaptcha(ctx *gin.Context) interface{} {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.IfErr(err)

	return gin.H{"captcha_id": id, "captcha_image": b64s}
}

func (c *VerifyCodeController) SendUsingPhone(ctx *gin.Context) interface{} {

	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if result := requests.Validate(ctx, &request, requests.VerifyCodePhone); result != nil {
		return result
	}

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(ctx, "发送短信失败~")
	}

	return true
}

// SendUsingEmail 发送 Email 验证码
func (c *VerifyCodeController) SendUsingEmail(ctx *gin.Context) interface{} {

	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if result := requests.Validate(ctx, &request, requests.VerifyCodeEmail); result != nil {
		return result
	}

	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(ctx, "发送 Email 验证码失败~")
	}

	return true
}

func (c *VerifyCodeController) Route(group *bingo.Group) {
	group.Group("auth", func(group *bingo.Group) {
		group.POST("verify-codes/captcha", c.showCaptcha)
		group.POST("verify-codes/phone", c.SendUsingPhone)
		group.POST("verify-codes/email", c.SendUsingEmail)
	})
}
