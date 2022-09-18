package repository

import "bingo-example/domain/entity/user"

// UserRepo 会员仓储
type UserRepo interface {
	Get(user *user.User) error
}
