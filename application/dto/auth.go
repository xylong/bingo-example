package dto

type (
	// LoginParam Login 账号密码登录
	LoginParam struct {
		Phone    string `json:"phone" form:"phone" binding:"required,phone"`
		Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
	}
)
