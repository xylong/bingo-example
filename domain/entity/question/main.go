package question

import "time"

// Question 题目
type Question struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`
	Code string `json:"code" gorm:"type:char(5);not null;comment:课程编码"`
	Question string `json:"question" gorm:"type:varchar(255);not null;comment:题目"`
	QuestionType uint8 `json:"question_type" gorm:"type:varchar(255);not null;comment:题目类型(1单选 2解释 3判断 4简答 5综合)"`
	QuestionImg string `json:"question_img" gorm:"type:varchar(255);default:null;comment:题目图片"`
	Day time.Time `json:"day" gorm:"type:date;default:null;comment:日期"`
}