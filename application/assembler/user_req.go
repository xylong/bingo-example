package assembler

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
)

// UserReq 用户请求
type UserReq struct{}

// Register2User 注册参数转用户实体
func (r *UserReq) Register2User(param *dto.RegisterParam) *user.User {
	return user.New(user.WithPhone(param.Phone))
}

func (r *UserReq) Register2Profile(param *dto.RegisterParam) *profile.Profile {
	return profile.New(profile.WithPassword(param.Password))
}

func (r *UserReq) Login2User(param *dto.LoginParam) *user.User {
	return user.New(user.WithPhone(param.Phone))
}
