package repository

import "bingo-example/domain/entity/question"

// QuestionRepo 题库仓储
type QuestionRepo interface {
	Create(*question.Question) error
}
