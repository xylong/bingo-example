package aggregate

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
)

// Member 会员
type Member struct {
	// User 根实体，🆔是聚合的主标识符
	User *user.User

	// 用户信息
	Profile *profile.Profile

	// Logs 登录日志
	Logs []*entity.LoginLog
}

func (m *Member) Builder(u *user.User) *MemberBuilder {
	return NewMemberBuilder(u)
}

func (m *Member) Create() {

}

func (m *Member) GetMembers() {

}
