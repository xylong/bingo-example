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

	// CountRegister 统计最近注册人数
	CountRegister(interface{}, string) (interface{}, error)
}
