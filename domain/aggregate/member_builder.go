package aggregate

import (
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
	"bingo-example/domain/repository"
	"bingo-example/infrastructure/dao/g"
	"gorm.io/gorm"
)

// MemberBuilder 会员构建器
// 建造者模式
type MemberBuilder struct {
	user    *user.User
	profile *profile.Profile
	logs    []*entity.LoginLog

	userRepo    repository.IUserRepo
	profileRepo repository.IProfileRepo
}

func (m *MemberBuilder) SetLogs(logs []*entity.LoginLog) *MemberBuilder {
	if logs != nil {
		m.logs = logs
	}

	return m
}

func (m *MemberBuilder) SetProfile(profile *profile.Profile) *MemberBuilder {
	if profile != nil {
		m.profile = profile
	}

	return m
}

func (m *MemberBuilder) SetUserRepo(db *gorm.DB) *MemberBuilder {
	m.userRepo = g.NewUserRepo(db)
	return m
}

func (m *MemberBuilder) SetProfileRepo(db *gorm.DB) *MemberBuilder {
	m.profileRepo = g.NewProfileRepo(db)
	return m
}

// Build 构建聚合实体
func (m *MemberBuilder) Build() *Member {
	member := &Member{User: m.user}

	if m.profile != nil {
		member.Profile = m.profile
	}

	if m.userRepo != nil {
		member.UserRepo = m.userRepo
	}

	if m.profileRepo != nil {
		member.ProfileRepo = m.profileRepo
	}

	return member
}

func NewMemberBuilder(user *user.User) *MemberBuilder {
	return &MemberBuilder{user: user}
}
