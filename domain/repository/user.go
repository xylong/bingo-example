package repository

import (
	"bingo-example/domain/entity/user"
	"gorm.io/gorm"
)

// IUserRepo 会员仓储
type IUserRepo interface {
	// Create 创建用户
	Create(*user.User) error

	// Get 根据查询条件获取用户
	Get(...func(db *gorm.DB) *gorm.DB) ([]*user.User, error)

	// GetCount 查询并统计用户数
	GetCount(...func(db *gorm.DB) *gorm.DB) (int64, []*user.User, error)

	// GetOne 查找单个用户
	GetOne(*user.User, ...func(*gorm.DB) *gorm.DB) error

	// CountRegister 统计最近注册人数
	CountRegister(interface{}, string) (interface{}, error)

	// IsPhoneExist 判断手机号是否存在
	IsPhoneExist(string) bool

	// IsEmailExist 判断邮箱是否已存在
	IsEmailExist(string) bool
}
