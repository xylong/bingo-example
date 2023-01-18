package aggregate

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
	"bingo-example/domain/entity/userlog"
	"bingo-example/domain/repository"
	"gorm.io/gorm"
)

// Member 会员
type Member struct {
	// User 根实体，🆔是聚合的主标识符
	User *user.User

	// 用户信息
	Profile *profile.Profile

	// Logs 登录日志
	Logs []*userlog.UserLog

	UserRepo    repository.IUserRepo
	ProfileRepo repository.IProfileRepo
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

func (m *Member) Take(with map[string][]string) error {
	return m.UserRepo.GetOne(m.User, entity.With(with))
}

// Get 获取用户
func (m *Member) Get(request *dto.UserRequest) (int64, []*user.User, error) {
	scopes := []func(db *gorm.DB) *gorm.DB{
		entity.Select("id", "nickname", "avatar", "phone", "email", "created_at"),
		entity.Paginate(request.Page, request.PageSize),
		entity.Order("id desc"),
	}

	{
		if request.Nickname != "" {
			scopes = append(scopes, m.User.NickNameCompare(request.Nickname, entity.Like))
		}

		if request.Phone != "" {
			scopes = append(scopes, m.User.PhoneCompare(request.Phone, entity.Equal))
		}

		if request.Email != "" {
			scopes = append(scopes, m.User.EmailCompare(request.Email, entity.Equal))
		}
	}

	return m.UserRepo.GetCount(scopes...)
}
