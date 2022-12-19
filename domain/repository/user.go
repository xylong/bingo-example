package repository

import "bingo-example/domain/entity/user"

// IUserRepo 会员仓储
type IUserRepo interface {
	Create(*user.User) error
	Get(*user.User, ...string) error
}
