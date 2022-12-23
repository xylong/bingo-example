package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/xylong/bingo"
	"strconv"
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

func (c *BookController) show(ctx *bingo.Context) interface{} {
	return c.Service.GetByID(ctx.Param("id"))
}

func (c *BookController) search(ctx *bingo.Context) interface{} {
	return c.Service.Search(
		ctx.Binding(ctx.ShouldBind, &dto.BookSearchParam{}).
			Unwrap().(*dto.BookSearchParam))
}

func (c *BookController) press(ctx *bingo.Context) interface{} {
	return c.Service.GetPress()
}

func (c *BookController) create(ctx *bingo.Context) interface{} {
	param := &dto.BookStoreParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return gin.H{"msg": err.Error()}
	}

	return c.Service.Create(param)
}

func (c *BookController) update(ctx *bingo.Context) interface{} {
	param := &dto.BookStoreParam{}
	if err := ctx.ShouldBind(param); err != nil {
		return gin.H{"msg": err.Error()}
	}

	id := ctx.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		return gin.H{"msg": err.Error()}
	}

	return c.Service.Update(_id, param)
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
}
