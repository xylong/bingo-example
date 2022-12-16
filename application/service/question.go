package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/domain/aggregate"
	"go.uber.org/zap"
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
	if err := new(aggregate.QuestionBank).
		Builder(s.Req.Param2QuestionModel(param)).
		SetQuestionRepo(s.DB).Build().
		Create(); err != nil {
		zap.L().Error("create question", zap.Error(err), zap.Any("param", param))
		return -1, "题库创建失败", nil
	}

	return 0, "", nil
}
