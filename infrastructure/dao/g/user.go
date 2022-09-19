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

func (d *UserDao) Get(user *user.User, with ...string) error {
	if len(with) > 0 {
		for _, s := range with {
			d.db = d.db.Preload(s)
		}
	}

	err := d.db.Where(user).First(user).Error
	if err != nil {
		return err
	}

	return nil
}
