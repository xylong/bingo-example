package repository

import "bingo-example/domain/entity/user"

// UserRepo 会员仓储
type UserRepo interface {
	Create(*user.User) error
	Get(user *user.User) error
}
