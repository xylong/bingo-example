package v1

import (
	"bingo-example/application/dto"
	"bingo-example/application/service"
	"github.com/xylong/bingo"
)

func init() {
	registerCtrl(NewQuestionController())
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

func (c *QuestionController) create(ctx *bingo.Context) (int, string, interface{}) {
	return c.Service.Create(
		ctx.Binding(ctx.ShouldBind, &dto.QuestionParam{}).
			Unwrap().(*dto.QuestionParam))
}

func (c *QuestionController) singleChoice(ctx *bingo.Context) string {
	return "单选题"
}

func (c *QuestionController) Route(group *bingo.Group) {
	group.POST("questions", c.create)
}
