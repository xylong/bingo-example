package v1

import (
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

func (c *QuestionController) singleChoice(ctx *bingo.Context) string {
	return "单选题"
}

func (c *QuestionController) Route(group *bingo.Group)  {
	group.GET("single-choice",c.singleChoice)
}