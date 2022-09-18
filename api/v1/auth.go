package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewAuthCtrl())
}

type AuthCtrl struct {
	Service *service.JwtService `inject:"-"`
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{}
}

func (c *AuthCtrl) Name() string {
	return "AuthCtrl"
}

func (c *AuthCtrl) login(ctx *bingo.Context) string {
	return c.Service.Login(
		ctx.Binding(ctx.ShouldBind, &dto.LoginParam{}).
			Unwrap().(*dto.LoginParam))

}

func (c *AuthCtrl) Route(group *bingo.Group) {
	group.POST("login", c.login)
}
