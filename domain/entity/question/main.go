package question

import (
	"bingo-example/domain/entity/answer"
	"time"
)

const (
	SingleChoice   = iota + 1 // 单选
	MultipleChoice            // 多选
	Judgment                  // 判断
	ShortAnswer               // 简单
	Comprehensive             // 综合
)

// Question 题目
type Question struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	Code         string    `json:"code" gorm:"type:char(5);not null;comment:课程编码"`
	Question     string    `json:"question" gorm:"type:varchar(255);not null;comment:题目"`
	QuestionType uint8     `json:"question_type" gorm:"type:varchar(255);not null;comment:题目类型(1单选 2解释 3判断 4简答 5综合)"`
	QuestionImg  string    `json:"question_img" gorm:"type:varchar(255);default:null;comment:题目图片"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_a"`

	// 关联
	Answer []*answer.Answer // 答案
}

func New(attributes ...Attribute) *Question {
	q := &Question{}
	Attributes(attributes).Apply(q)
	return q
}
