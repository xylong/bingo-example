package aggregate

import (
	"bingo-example/domain/entity/answer"
	"bingo-example/domain/entity/question"
	"bingo-example/domain/repository"
	"bingo-example/infrastructure/dao/g"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// questionBankBuilder 题库构造器
type questionBankBuilder struct {
	question *question.Question
	answers  []*answer.Answer

	questionRepo repository.QuestionRepo
}

func newQuestionBankBuilder(question *question.Question) *questionBankBuilder {
	return &questionBankBuilder{question: question}
}

func (b *questionBankBuilder) SetAnswer(answers []*answer.Answer) *questionBankBuilder {
	if answers != nil && len(answers) > 0 {
		b.answers = answers
	}

	return b
}

func (b *questionBankBuilder) SetQuestionRepo(db *gorm.DB) *questionBankBuilder {
	if db != nil {
		b.questionRepo = g.NewQuestionDao(db)
	} else {
		zap.L().Warn("question repo is nil")
	}

	return b
}

func (b *questionBankBuilder) Build() *QuestionBank {
	qb := &QuestionBank{Question: b.question}

	if b.answers != nil {
		qb.Answers = b.answers
	}

	if b.questionRepo != nil {
		qb.QuestionRepo = b.questionRepo
	}

	return qb
}
