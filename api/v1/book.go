package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/gin-gonic/gin"
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
	c.Service.BatchImport(ctx)
}

func (c *BookController) index(ctx *bingo.Context) {
	schema := c.Service.GraphSchema(ctx)

	handler.New(&handler.Config{
		Schema: &schema,
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *BookController) show(ctx *bingo.Context) interface{} {
	return c.Service.GetByID(ctx, ctx.Param("id"))
}

func (c *BookController) search(ctx *bingo.Context) interface{} {
	return c.Service.Search(ctx, ctx.Binding(ctx.ShouldBind, &dto.BookSearchParam{}).Unwrap().(*dto.BookSearchParam))
}

func (c *BookController) press(ctx *bingo.Context) interface{} {
	return c.Service.GetPress(ctx)
}

func (c *BookController) create(ctx *bingo.Context) interface{} {
	param := &dto.BookStoreParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return gin.H{"msg": err.Error()}
	}

	return c.Service.Create(ctx, param)
}

func (c *BookController) update(ctx *bingo.Context) interface{} {
	url := &dto.BookUrlRequest{}
	if err := ctx.ShouldBindUri(url); err != nil {
		return gin.H{"msg": err.Error()}
	}

	param := &dto.BookStoreParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return gin.H{"msg": err.Error()}
	}

	return c.Service.Update(ctx, url, param)
}

func (c *BookController) delete(ctx *bingo.Context) interface{} {
	if err := c.Service.Delete(ctx, ctx.Binding(ctx.ShouldBindUri, &dto.BookUrlRequest{}).Unwrap().(*dto.BookUrlRequest)); err != nil {
		return gin.H{"code": 10000, "msg": err.Error(), "data": nil}
	}

	return gin.H{"code": 0, "msg": "", "data": nil}
}

func (c *BookController) Name() string {
	return "BookController"
}

func (c *BookController) Route(group *bingo.Group) {
	group.GET("import", c.import2es)
	group.GET("books", c.search)
	group.GET("books/:id", c.show)
	group.GET("presses", c.press)
	group.POST("book", c.index)
	group.POST("books", c.create)
	group.PUT("books/:id", c.update)
	group.DELETE("books/:id", c.delete)
}
