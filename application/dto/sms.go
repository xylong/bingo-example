package dto

// SmsParam 短信验证
type SmsParam struct {
	Phone string `json:"phone" form:"phone" binding:"required,phone"`
	Code  int    `json:"code" form:"code" binding:"sms"`
}
