package assembler

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity/user"
)

// UserRep 用户响应
type UserRep struct{}

func (r *UserRep) User2Profile(user *user.User) *dto.Profile {
	gender := ""
	switch user.Profile.Gender {
	case 0:
		gender = "女"
	case 1:
		gender = "男"
	default:
		gender = "保密"
	}

	return &dto.Profile{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Phone:     user.Phone,
		Email:     user.Email,
		Birthday:  user.Profile.Birth(),
		Gender:    gender,
		Level:     user.Profile.Level,
		Signature: user.Profile.Signature,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (r *UserRep) SimpleList(users []*user.User) []*dto.SimpleUser {
	if users == nil || len(users) == 0 {
		return nil
	}

	var list []*dto.SimpleUser
	for _, u := range users {
		list = append(list, &dto.SimpleUser{
			ID:     u.ID,
			Avatar: u.Avatar,
			Phone:  u.Phone,
			Email:  u.Email,
			//Gender:    u.Profile.Gender,
			//Level:     u.Profile.Level,
			Signature: u.Nickname,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return list
}
