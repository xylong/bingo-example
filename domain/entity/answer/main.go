package answer

// Answer 答案
type Answer struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`
	QuestionID int `json:"question_id" gorm:"type:int(11);not null;comment:问题id"`
	Answer string `json:"answer" gorm:"type:text;not null;comment:答案"`
	Correct uint8 `json:"correct" gorm:"type:tinyint(1);default:0;comment:0正确 1错误"`
	Other string `json:"other" gorm:"type:varchar(100);default:null;comment:额外补充"`
	Img string `json:"img" gorm:"type:varchar(255);default:null;comment:图片"`
}
