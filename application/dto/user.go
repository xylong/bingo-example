package dto

type (
	// UserRequest 用户查询请求
	UserRequest struct {
		*Pagination
		Nickname string `form:"nickname" json:"nickname" binding:"omitempty"`
		Phone    string `form:"phone" json:"phone" binding:"omitempty,phone"`
		Email    string `form:"email" json:"email" binding:"omitempty,email"`
		Birthday string `form:"birthday" json:"birthday" binding:"omitempty,date"`          // 生日
		Gender   int8   `form:"gender" json:"gender" binding:"omitempty,oneof=-1 0 1"`      // 性别
		Level    uint8  `form:"level" json:"level" binding:"omitempty,oneof=0 1 2 3 4 5 6"` // 等级
	}
)

type (
	// SimpleUserList 简单用户列表
	SimpleUserList struct {
		Total int64         `json:"total"`
		List  []*SimpleUser `json:"list"`
	}

	// SimpleUser 简洁用户信息
	SimpleUser struct {
		ID        int    `json:"id"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Gender    int8   `json:"gender"`
		Level     uint8  `json:"level"`
		Signature string `json:"signature"`
		CreatedAt string `json:"created_at"`
	}

	// Profile 个人信息
	Profile struct {
		ID        int    `json:"id"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Birthday  string `json:"birthday"`
		Gender    string `json:"gender"`
		Level     uint8  `json:"level"`
		Signature string `json:"signature"`
		CreatedAt string `json:"created_at"`
	}
)
