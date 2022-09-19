package aggregate

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
	"bingo-example/domain/repository"
)

// Member 会员
type Member struct {
	// User 根实体，🆔是聚合的主标识符
	User *user.User

	// 用户信息
	Profile *profile.Profile

	// Logs 登录日志
	Logs []*entity.LoginLog

	UserRepo    repository.UserRepo
	ProfileRepo repository.ProfileRepo
}

func (m *Member) Builder(u *user.User) *MemberBuilder {
	return NewMemberBuilder(u)
}

func (m *Member) Create() error {
	if err := m.UserRepo.Create(m.User); err != nil {
		return err
	}

	m.Profile.UserID = m.User.ID
	if err := m.ProfileRepo.Create(m.Profile); err != nil {
		return err
	}

	return nil
}

func (m *Member) GetMembers() {

}
