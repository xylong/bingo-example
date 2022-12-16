package aggregate

import (
	"bingo-example/domain/entity/answer"
	"bingo-example/domain/entity/question"
	"bingo-example/domain/repository"
)

// QuestionBank 题库
type QuestionBank struct {
	// 题干
	Question *question.Question

	// 答案
	Answers []*answer.Answer

	QuestionRepo repository.QuestionRepo
}

func (b *QuestionBank) Builder(question *question.Question) *questionBankBuilder {
	return newQuestionBankBuilder(question)
}

// Create 创建
func (b *QuestionBank) Create() error {
	if err := b.QuestionRepo.Create(b.Question); err != nil {
		return err
	}

	return nil
}
