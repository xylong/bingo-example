package assembler

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
	"golang.org/x/crypto/bcrypt"
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

// EncryptPassword 加密密码
func (r *UserReq) EncryptPassword(password string) string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd)
}
