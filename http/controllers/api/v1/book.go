package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"bingo-example/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewBookController())
}

type BookController struct {
	Service *service.BookService `inject:"-"`
}

func NewBookController() *BookController {
	return &BookController{}
}

func (c *BookController) index(ctx *gin.Context) {
	schema := c.Service.GraphSchema(ctx)

	handler.New(&handler.Config{
		Schema: &schema,
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *BookController) show(ctx *gin.Context) interface{} {
	return c.Service.GetByID(ctx, ctx.Param("id"))
}

func (c *BookController) search(ctx *gin.Context) interface{} {
	param := &dto.BookSearchParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return response.Error(ctx, err)
	}

	return c.Service.Search(ctx, param)
}

func (c *BookController) press(ctx *gin.Context) interface{} {
	return c.Service.GetPress(ctx)
}

func (c *BookController) create(ctx *gin.Context) interface{} {
	param := &dto.BookStoreParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return gin.H{"msg": err.Error()}
	}

	return c.Service.Create(ctx, param)
}

func (c *BookController) update(ctx *gin.Context) interface{} {
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

func (c *BookController) delete(ctx *gin.Context) interface{} {
	//if err := c.Service.Delete(ctx, ctx.Binding(ctx.ShouldBindUri, &dto.BookUrlRequest{}).Unwrap().(*dto.BookUrlRequest)); err != nil {
	//	return gin.H{"code": 10000, "msg": err.Error(), "data": nil}
	//}

	return gin.H{"code": 0, "msg": "", "data": nil}
}

func (c *BookController) Name() string {
	return "BookController"
}

func (c *BookController) Route(group *bingo.Group) {
	group.GET("books", c.search)
	group.GET("books/:id", c.show)
	group.GET("presses", c.press)
	group.POST("book", c.index)
	group.POST("books", c.create)
	group.PUT("books/:id", c.update)
	group.DELETE("books/:id", c.delete)
}
