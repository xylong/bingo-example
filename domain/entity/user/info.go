package user

// Info 用户信息
type Info struct {
	Avatar   string `gorm:"type:varchar(100);comment:头像" json:"avatar" bson:"avatar"`
	Nickname string `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname" bson:"nickname"`
	Phone    string `gorm:"type:char(11);uniqueIndex;comment:手机号" json:"phone" bson:"phone"`
	Email    string `gorm:"type:varchar(50);default:null;uniqueIndex;comment:邮件" json:"email" bson:"email"`
}

func NewInfo() *Info {
	return &Info{}
}
