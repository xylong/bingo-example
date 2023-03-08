package user

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID int `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`

	*ThirdParty `gorm:"embedded"` // 第三方信息
	*Info       `gorm:"embedded"` // 	基本信息

	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;" bson:"deleted_at"`

	// 关联
	Profile *profile.Profile // has one
}

func New(attr ...Attr) *User {
	user := &User{ThirdParty: NewThirdParty(), Info: NewInfo()}
	Attrs(attr).Apply(user)
	return user
}

// NickNameCompare 比较名称
func (u *User) NickNameCompare(name string, operator int) entity.Scope {
	return func(db *gorm.DB) *gorm.DB {
		switch operator {
		case entity.NotEqual:
			return db.Where("nickname<>?", name)
		case entity.Like:
			return db.Where("nickname like ?", "%"+name+"%")
		case entity.NotLike:
			return db.Where("nickname not like ?", "%"+name+"%")
		default:
			return db.Where("nickname=?", name)
		}
	}
}

// PhoneCompare 比较手机号
func (u *User) PhoneCompare(phone string, operator int) entity.Scope {
	return func(db *gorm.DB) *gorm.DB {
		switch operator {
		case entity.NotEqual:
			return db.Where("phone<>?", phone)
		default:
			return db.Where("phone=?", phone)
		}
	}
}

func (u *User) EmailCompare(email string, operator int) entity.Scope {
	return func(db *gorm.DB) *gorm.DB {
		switch operator {
		case entity.NotEqual:
			return db.Where("email<>?", email)
		default:
			return db.Where("email=?", email)
		}
	}
}
