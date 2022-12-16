package g

import (
	"bingo-example/domain/entity/question"
	"gorm.io/gorm"
)

type QuestionDao struct {
	db *gorm.DB
}

func NewQuestionDao(db *gorm.DB) *QuestionDao {
	return &QuestionDao{db: db}
}

// Create 创建
func (d QuestionDao) Create(question *question.Question) error {
	return d.db.Create(question).Error
}
