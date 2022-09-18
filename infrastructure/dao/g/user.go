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

func (r *UserRepo) Get(user *user.User) error {
	err := r.db.First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
