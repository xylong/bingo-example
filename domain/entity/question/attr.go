package question

import (
	"bingo-example/domain/entity/answer"
)

type (
	// Attribute 属性
	Attribute  func(*Question)
	Attributes []Attribute
)

// Apply 应用属性函数
func (a Attributes) Apply(question *Question) {
	for _, attribute := range a {
		attribute(question)
	}
}

func WithID(id int) Attribute {
	return func(question *Question) {
		if id > 0 {
			question.ID = id
		}
	}
}

func WithCode(code string) Attribute {
	return func(question *Question) {
		if code != "" {
			question.Code = code
		}
	}
}

func WithQuestion(question string) Attribute {
	return func(q *Question) {
		if question != "" {
			q.Question = question
		}
	}
}

func WithQuestionType(questionType uint8) Attribute {
	return func(question *Question) {
		if questionType == SingleChoice ||
			questionType == MultipleChoice ||
			questionType == Judgment ||
			questionType == ShortAnswer ||
			questionType == Comprehensive {
			question.QuestionType = questionType
		}
	}
}

func WithQuestionImg(img string) Attribute {
	return func(question *Question) {
		if img != "" {
			question.QuestionImg = img
		}
	}
}

func WithAnswer(answer []*answer.Answer) Attribute {
	return func(question *Question) {
		if answer != nil && len(answer) > 0 {
			question.Answer = answer
		}
	}
}
