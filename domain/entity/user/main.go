package user

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	Nickname string `json:"name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Gender   uint8  `json:"gender,omitempty"`

	*ThirdParty `gorm:"embedded"` // 第三方信息
	*Info       `gorm:"embedded"` // 	基本信息

	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;" bson:"deleted_at"`

	// 关联
	Profile *profile.Profile
}

func NewUser(attr ...entity.Attr) *User {
	user := &User{ThirdParty: NewThirdParty(), Info: NewInfo()}
	entity.Attrs(attr).Apply(user)
	return user
}

func WithID(id int) entity.Attr {
	return func(i interface{}) {
		if id > 0 {
			i.(*User).ID = id
		}
	}
}

func WithPhone(phone string) entity.Attr {
	return func(i interface{}) {
		if phone != "" {
			i.(*User).Phone = phone
		}
	}
}
