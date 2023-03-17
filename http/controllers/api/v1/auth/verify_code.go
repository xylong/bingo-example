package auth

import (
	v1 "bingo-example/http/controllers/api/v1"
	"bingo-example/pkg/captcha"
	"bingo-example/pkg/logger"
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

func (c *VerifyCodeController) show(ctx *gin.Context) interface{} {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.IfErr(err)

	return gin.H{"captcha_id": id, "captcha_image": b64s}
}

func (c *VerifyCodeController) Route(group *bingo.Group) {
	group.Group("auth", func(group *bingo.Group) {
		group.POST("verify-codes/captcha", c.show)
	})
}
