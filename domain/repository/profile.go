package repository

import "bingo-example/domain/entity/profile"

// IProfileRepo 用户信息
type IProfileRepo interface {
	Create(*profile.Profile) error
}
