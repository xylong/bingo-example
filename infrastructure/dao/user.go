package dao

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

func (r *UserRepo) Create(user *user.User) error {
	return r.db.Create(user).Error
}

// Get 获取用户
func (r *UserRepo) Get(scopes ...func(db *gorm.DB) *gorm.DB) ([]*user.User, error) {
	var users []*user.User

	if err := r.db.Scopes(scopes...).Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// GetCount 查询并统计用户数
func (r *UserRepo) GetCount(scopes ...func(db *gorm.DB) *gorm.DB) (total int64, users []*user.User, err error) {
	err = r.db.Scopes(scopes...).Find(&users).Limit(-1).Offset(-1).Count(&total).Error
	return
}

// CountRegister 统计注册数
func (r *UserRepo) CountRegister(result interface{}) (interface{}, error) {
	err := r.db.Model(user.New()).
		Select("date(created_at) as date, count(*) as total").
		Group("date").Having("total > ?", 0).
		Scan(&result).Error

	return result, err
}
