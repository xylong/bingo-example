package user

import (
	"bingo-example/domain/entity"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	Name     string `bson:"name" json:"name,omitempty"`
	Phone    string `bson:"phone" json:"phone,omitempty"`
	Birthday string `bson:"birthday" json:"birthday,omitempty"`
	Gender   uint8  `bson:"gender" json:"gender,omitempty"`

	*ThirdParty `gorm:"embedded"` // 第三方信息
	*Info       `gorm:"embedded"` // 	基本信息

	CreatedAt *entity.DateAt `json:"created_at" bson:"created_at"`
	UpdatedAt *entity.DateAt `json:"updated_at" bson:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;" bson:"deleted_at"`
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
