package g

import (
	"bingo-example/domain/entity/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (d *UserRepo) Create(user *user.User) error {
	return d.db.Create(user).Error
}

func (d *UserRepo) Get(user *user.User, with ...string) error {
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
