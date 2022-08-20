package aggregate

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/user"
)

// Member 会员
type Member struct {
	// User 根实体，🆔是聚合的主标识符
	User *user.User

	// Logs 登录日志
	Logs []*entity.LoginLog
}

func NewMember() *Member {
	return &Member{}
}
