package profile

import (
	"bingo-example/infrastructure/util"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// Profile 用户信息
type Profile struct {
	ID        int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	UserID    int       `gorm:"type:int(11);not null;uniqueIndex;comment:用户🆔" json:"user_id"`
	Password  string    `gorm:"type:varchar(100);comment:密码" json:"-"`
	Salt      string    `gorm:"type:char(6);comment:盐" json:"salt"`
	Birthday  time.Time `gorm:"type:date;default:null;comment:出生日期" json:"birthday"`
	Gender    int8      `gorm:"type:tinyint(1);default:-1;comment:-1保密 0女 1男" json:"gender"`
	Level     uint8     `gorm:"type:tinyint(1);default:0;comment:等级" json:"level"`
	Signature string    `gorm:"type:varchar(255);comment=个性签名" json:"signature"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(attr ...Attr) *Profile {
	p := &Profile{}
	Attrs(attr).Apply(p)
	return p
}

func (p *Profile) BeforeCreate(db *gorm.DB) error {
	if p == nil {
		return errors.New("can't save invalid data")
	}

	p.Salt = util.RandString(6)
	pwd, _ := bcrypt.GenerateFromPassword([]byte(util.Md5(p.Password)+p.Salt), bcrypt.DefaultCost)
	p.Password = string(pwd)

	return nil
}

func (p *Profile) Birth() string {
	if !p.Birthday.IsZero() {
		return p.Birthday.Format("2006-01-02")
	}

	return ""
}

// VerifyPassword 校验密码
func (p *Profile) VerifyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(util.Md5(password)+p.Salt)) == nil
}

// GenderString 中文性别
func (p *Profile) GenderString() string {
	var gender string

	switch p.Gender {
	case 0:
		gender = "女"
	case 1:
		gender = "男"
	default:
		gender = "保密"
	}

	return gender
}
