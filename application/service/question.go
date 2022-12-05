package service

import "bingo-example/application/assembler"

// QuestionService 题库
type QuestionService struct {
	Req *assembler.QuestionReq
	Rep assembler.QuestionRep
}

// SingleChoice 单选题
func (s *QuestionService) SingleChoice()  {

}