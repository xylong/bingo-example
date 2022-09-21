package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/middleware"
	"bingo-example/application/service"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewUserCtrl())
}

type UserCtrl struct {
	Service *service.UserService `inject:"-"`
	Jwt     *service.JwtService  `inject:"-"`
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (c *UserCtrl) Name() string {
	return "UserCtrl"
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
	return c.Service.Profile(ctx.GetInt("user_id"))
}

func (c *UserCtrl) Route(group *bingo.Group) {
	group.POST("register", c.register)
	group.POST("login", c.login)

	group.Group("", func(group *bingo.Group) {
		group.GET("me", c.me)
	}, middleware.NewAuthentication())
}
