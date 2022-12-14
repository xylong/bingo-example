package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"gorm.io/gorm"
)

// QuestionService 题库
type QuestionService struct {
	Req *assembler.QuestionReq
	Rep assembler.QuestionRep

	DB *gorm.DB `inject:"-"`
}

// Create 创建
func (s *QuestionService) Create(param *dto.QuestionParam) (int, string, interface{}) {
	return 0, "", s.Req.Param2QuestionModel(param)
}
