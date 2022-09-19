package g

import (
	"bingo-example/domain/entity/user"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Create(user *user.User) error {
	return d.db.Create(user).Error
}

func (d *UserDao) Get(user *user.User) error {
	err := d.db.First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
