package repository

import "bingo-example/domain/entity/question"

// IQuestionRepo 题库仓储
type IQuestionRepo interface {
	Create(*question.Question) error
}
