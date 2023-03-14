package v1

import (
	"bingo-example/application/service"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewQuestionController())
}

type QuestionController struct {
	Service *service.QuestionService `inject:"-"`
}

func NewQuestionController() *QuestionController {
	return &QuestionController{}
}

func (c *QuestionController) Name() string {
	return "questions"
}

func (c *QuestionController) create(ctx *gin.Context) (int, string, interface{}) {
	return 0, "", nil
	//return c.Service.Create(
	//	ctx.Binding(ctx.ShouldBind, &dto.QuestionParam{}).
	//		Unwrap().(*dto.QuestionParam))
}

func (c *QuestionController) singleChoice(ctx *gin.Context) string {
	return "单选题"
}

func (c *QuestionController) Route(group *bingo.Group) {
	group.POST("questions", c.create)
}
