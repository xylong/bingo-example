package dto

type (
	// Profile 个人信息
	Profile struct {
		ID        int    `json:"id"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Birthday  string `json:"birthday"`
		Gender    string `json:"gender"`
		Level     int8   `json:"level"`
		Signature string `json:"signature"`
		CreatedAt string `json:"created_at"`
	}
)
