package g

import (
	"bingo-example/domain/entity/question"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

// Create 创建
func (d QuestionRepo) Create(question *question.Question) error {
	return d.db.Create(question).Error
}
