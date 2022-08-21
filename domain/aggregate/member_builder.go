package aggregate

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
)

// MemberBuilder 会员构建器
// 建造者模式
type MemberBuilder struct {
	user    *user.User
	profile *profile.Profile
	logs    []*entity.LoginLog
}

func (m *MemberBuilder) SetProfile(profile *profile.Profile) *MemberBuilder {
	m.profile = profile
	return m
}

// Build 构建聚合实体
func (m *MemberBuilder) Build() *Member {
	member := &Member{User: m.user}

	if m.profile != nil {
		member.Profile = m.profile
	}

	return member
}

func NewMemberBuilder(user *user.User) *MemberBuilder {
	return &MemberBuilder{user: user}
}
