package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/graphql-go/handler"
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

func (c *BookController) index(ctx *bingo.Context) {
	schema := c.Service.GraphSchema()

	handler.New(&handler.Config{
		Schema: &schema,
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *BookController) search(ctx *bingo.Context) interface{} {
	return c.Service.Search(
		ctx.Binding(ctx.ShouldBind, &dto.BookSearchParam{}).
			Unwrap().(*dto.BookSearchParam))
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
	group.GET("presses", c.press)
	group.POST("book", c.index)
}
