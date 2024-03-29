package dto

type (
	// RegisterParam 注册
	RegisterParam struct {
		Phone    string `json:"phone" form:"phone" binding:"required,phone"` // 手机号
		Password string `json:"password" form:"password" binding:"required,min=6,max=18"	// 密码`
		Code     int    `json:"code" form:"code" binding:"required"` // 短信码
	}

	// LoginParam Login 账号密码登录
	LoginParam struct {
		Phone    string `json:"phone" form:"phone" binding:"required,phone"`
		Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
	}
)
