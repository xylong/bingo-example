package v1

import (
	"bingo-example/application/service"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewFruitController())
}

type FruitController struct {
	Service *service.FruitService `inject:"-"`
}

func NewFruitController() *FruitController {
	return &FruitController{}
}

func (c *FruitController) Name() string {
	return "FruitController"
}

func (c *FruitController) Route(group *bingo.Group) {
	group.GET("top", c.top)
}

func (c *FruitController) top(ctx *gin.Context) interface{} {
	return c.Service.Top()
}
