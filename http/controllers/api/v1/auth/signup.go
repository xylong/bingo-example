package auth

import (
	v1 "bingo-example/http/controllers/api/v1"
	"bingo-example/infrastructure/dao"
	"bingo-example/pkg/database"
	"bingo-example/pkg/response"
	"bingo-example/requests"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/log"
)

func init() {
	v1.RegisterCtrl(NewSignupController())
}

// SignupController 注册控制器
type SignupController struct {
	*v1.BaseApiController
}

func NewSignupController() *SignupController {
	return &SignupController{}
}

func (c *SignupController) Name() string {
	return "signup"
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) interface{} {
	request := requests.SignupPhoneExistRequest{}
	if result := requests.Validate(ctx, &request, requests.SignupPhoneExist); result != nil {
		return result
	}

	log.Dump(request)

	return response.Data(dao.NewUserRepo(database.DB()).IsPhoneExist(request.Phone))
}

func (c *SignupController) IsEmailExist(ctx *gin.Context) interface{} {
	request := requests.SignupEmailExistRequest{}
	if result := requests.Validate(ctx, &request, requests.SignupEmailExist); result != nil {
		return result
	}

	return response.Data(dao.NewUserRepo(database.DB()).IsEmailExist(request.Email))
}

func (c *SignupController) Route(group *bingo.Group) {
	group.Group("auth", func(group *bingo.Group) {
		group.POST("signup/phone/exist", c.IsPhoneExist)
		group.POST("signup/email/exist", c.IsEmailExist)
	})
}
