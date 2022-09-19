package repository

import "bingo-example/domain/entity/profile"

// ProfileRepo 用户信息
type ProfileRepo interface {
	Create(*profile.Profile) error
}
