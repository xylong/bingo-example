package answer

import "time"

// Answer 答案
type Answer struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	QuestionID int       `json:"question_id" gorm:"type:int(11);not null;comment:问题id"`
	Content    string    `json:"content" gorm:"type:text;not null;comment:内容"`
	IsCorrect  bool      `json:"is_correct" gorm:"type:tinyint(1);default:true;comment:是否正确"`
	Other      string    `json:"other" gorm:"type:varchar(100);default:null;comment:额外补充"`
	Img        string    `json:"img" gorm:"type:varchar(255);default:null;comment:图片"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_a"`
}

func New(attributes ...Attribute) *Answer {
	a := &Answer{}
	Attributes(attributes).Apply(a)
	return a
}
