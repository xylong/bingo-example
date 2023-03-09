package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/middleware"
	"bingo-example/application/service"
	"bingo-example/constants"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

func init() {
	registerCtrl(NewUserCtrl())
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

func (c *UserCtrl) index(ctx *bingo.Context) interface{} {
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

func (c *UserCtrl) register(ctx *bingo.Context) (int, string, interface{}) {
	return c.Service.Register(
		ctx.Binding(ctx.ShouldBind, &dto.RegisterParam{}).
			Unwrap().(*dto.RegisterParam))
}

func (c *UserCtrl) login(ctx *bingo.Context) (int, string, interface{}) {
	return c.Service.Login(
		ctx.Binding(ctx.ShouldBind, &dto.LoginParam{}).
			Unwrap().(*dto.LoginParam))
}

func (c *UserCtrl) me(ctx *bingo.Context) (int, string, interface{}) {
	return c.Service.Profile(ctx.GetInt(constants.SessionID))
}

func (c *UserCtrl) countRegister(ctx *bingo.Context) interface{} {
	return c.Service.CountReg(ctx,
		ctx.Binding(ctx.ShouldBind, &dto.RegisterCountRequest{}).
			Unwrap().(*dto.RegisterCountRequest))
}

func (c *UserCtrl) Route(group *bingo.Group) {
	group.GET("users", c.index)
	group.POST("register", c.register)
	group.POST("login", c.login)
	group.GET("reg-count", c.countRegister)

	group.Group("", func(group *bingo.Group) {
		group.GET("me", c.me)
	}, middleware.NewAuthentication())
}
