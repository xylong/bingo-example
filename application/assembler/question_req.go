package assembler

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity/answer"
	"bingo-example/domain/entity/question"
)

type QuestionReq struct{}

func (r *QuestionReq) Param2QuestionModel(param *dto.QuestionParam) *question.Question {
	q := question.New(
		question.WithID(param.ID),
		question.WithCode(param.Code),
		question.WithQuestion(param.Question),
		question.WithQuestionType(param.QuestionType),
		question.WithQuestionImg(param.QuestionImg),
	)

	if len(param.Answer) > 0 {
		arr := make([]*answer.Answer, 0)
		for _, answerParam := range param.Answer {
			arr = append(arr, r.Param2AnswerModel(answerParam))
		}

		q.Answer = arr
	}

	return q
}

func (r *QuestionReq) Param2AnswerModel(param *dto.AnswerParam) *answer.Answer {
	return answer.New(
		answer.WithID(param.ID),
		answer.WithQuestionID(param.QuestionID),
		answer.WithContent(param.Content),
		answer.WithIsCorrect(param.IsCorrect),
		answer.WithOther(param.Other),
		answer.WithImg(param.Img),
	)
}
