package dto

type (
	// QuestionParam 题目
	QuestionParam struct {
		ID           int            `json:"id" form:"id" binding:"omitempty,gt=0"`
		Code         string         `json:"code" form:"code" binding:"required"`
		Question     string         `json:"question" form:"question" binding:"required"`
		QuestionType uint8          `json:"question_type" form:"question_type" binding:"required,oneof=1 2 3 4 5"`
		QuestionImg  string         `json:"question_img" form:"question_img" binding:"omitempty"`
		Answer       []*AnswerParam `json:"answer" form:"answer" binding:"required"`
	}

	// AnswerParam 答案
	AnswerParam struct {
		ID         int    `json:"id" form:"id" binding:"omitempty,gt=0"`
		QuestionID int    `json:"question_id" form:"question_id" binding:"omitempty,gt=0"`
		Content    string `json:"content" form:"content" binding:"required"`
		IsCorrect  bool   `json:"is_correct" form:"is_correct" binding:"omitempty"`
		Other      string `json:"other" form:"other" binding:"omitempty"`
		Img        string `json:"img" form:"img" binding:"omitempty"`
	}
)
