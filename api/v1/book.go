package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/xylong/bingo"
)

func init() {
	registerCtrl(NewBookController())
}

type BookController struct {
	Service *service.BookService `inject:"-"`
}

func NewBookController() *BookController {
	return &BookController{}
}

func (c *BookController) import2es(ctx *bingo.Context) {
	c.Service.BatchImport()
}

func (c *BookController) search(ctx *bingo.Context) interface{} {
	return c.Service.Search(
		ctx.Binding(ctx.ShouldBind, &dto.BookQuery{}).
			Unwrap().(*dto.BookQuery))
}

func (c *BookController) press(ctx *bingo.Context) interface{} {
	return c.Service.GetPress()
}

func (c *BookController) Name() string {
	return "BookController"
}

func (c *BookController) Route(group *bingo.Group) {
	group.GET("import", c.import2es)
	group.GET("books", c.search)
	group.GET("press", c.press)
}
