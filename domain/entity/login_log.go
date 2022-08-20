package entity

const (
	Login  = 1 // 登录
	Logout = 2 // 登出

	FailedLog  = 0 // 失败
	SucceedLog = 1 // 成功
)

// LoginLog 登录日志
type LoginLog struct {
	UserID    int    `bson:"user_id" json:"user_id"` // 用户🆔
	IP        string `bson:"ip" json:"ip"`           // ip地址
	Type      uint8  `bson:"type"`                   // 1登录 2登出
	Status    uint8  `bson:"status"`                 // 0失败 1成功
	CreatedAt string `bson:"created_at"`             // 创建时间
}

func NewLoginLog(attr ...Attr) *LoginLog {
	log := &LoginLog{}
	Attrs(attr).Apply(log)
	return log
}
