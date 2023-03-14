package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"bingo-example/constants"
	"bingo-example/constants/errors"
	"bingo-example/http/middlewares"
	"bingo-example/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"strconv"
)

func init() {
	RegisterCtrl(NewUserCtrl())
}

type UserCtrl struct {
	Service *service.UserService `inject:"-"`
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (c *UserCtrl) Name() string {
	return "UserCtrl"
}

func (c *UserCtrl) index(ctx *gin.Context) interface{} {
	req := &dto.UserRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		return err.Error()
	}

	if result, err := c.Service.Get(ctx, req); err != nil {
		return err.Error()
	} else {
		return gin.H{"code": 0, "msg": "", "data": result}
	}
}

// @Summary 用户详情
// @Description 手机号码注册
// @Tags 用户
// @Produce  json
// @Param id path int true "用户id"
// @Success 200 {object} dto.Profile "用户信息"
// @Failure 404 {object} any "{"code":404,"data":null,"msg":"未查询到结果"}"
// @Router /v1/users/{id} [get]
func (c *UserCtrl) show(ctx *gin.Context) interface{} {
	id, err := strconv.Atoi(ctx.Param("id"))
	{
		if err != nil {
			return response.BadRequest(err)
		}
		if id <= 0 {
			return response.BadRequest(fmt.Errorf("id必须大于0"))
		}
	}

	if data, err := c.Service.Profile(id); err != nil {
		return response.Error(ctx, err)
	} else {
		return response.Data(data)
	}
}

// @Summary 注册
// @Description 手机号码注册
// @Tags 用户
// @Produce  json
// @Param param body dto.RegisterParam  true "注册表单"
// @Success 200 {object} any "{"code":0,"data":{"access_token":"","refresh_token":""},"msg":"ok"}"
// @Failure 400 {object} any "{"code":400,"data":null,"msg":"参数错误"}"
// @Router /v1/register [post]
func (c *UserCtrl) register(ctx *gin.Context) interface{} {
	var param dto.RegisterParam
	if err := ctx.ShouldBind(&param); err != nil {
		return response.BadRequest(err)
	}

	return c.Service.Register(&param)
}

func (c *UserCtrl) login(ctx *gin.Context) (int, string, interface{}) {
	param := &dto.LoginParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return errors.ParamError.Int(), errors.ParamError.String(), nil
	}

	return c.Service.Login(param)
}

func (c *UserCtrl) resetPassword(ctx *gin.Context) interface{} {
	fmt.Println("reset")
	return true
}

// @Summary 个人信息
// @Description 登录人信息
// @Tags 用户
// @Security ApiKeyAuth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} dto.Profile "success"
// @Router /v1/me [get]
func (c *UserCtrl) me(ctx *gin.Context) interface{} {
	if data, err := c.Service.Profile(ctx.GetInt(constants.SessionID)); err != nil {
		return response.Error(ctx, err)
	} else {
		return response.Data(data)
	}
}

// @Summary 注册统计
// @Description 按月统计当月每天注册人数
// @Tags 用户
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param month query string true "Y-d"
// @success 200 {object} any{code=int,data=[]dto.RegisterCount,message=string} "结果按日期分组"
// @Router /v1/reg-count [get]
func (c *UserCtrl) countRegister(ctx *gin.Context) interface{} {
	var param *dto.RegisterCountRequest
	if err := ctx.ShouldBind(&param); err != nil {
		return response.BadRequest(err)
	}

	return response.Data(c.Service.CountReg(ctx, param))
}

func (c *UserCtrl) Route(group *bingo.Group) {
	group.GET("users", c.index)
	group.GET("users/:id", c.show)
	group.POST("register", c.register)
	group.POST("login", c.login)
	group.GET("reg-count", c.countRegister)

	group.Group("", func(group *bingo.Group) {
		group.GET("me", c.me)

		group.Group("reset", func(group *bingo.Group) {
			group.POST("password", c.resetPassword)
		}, middlewares.Request)

	}, middlewares.Auth)
}
