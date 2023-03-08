package assembler

import (
	"bingo-example/application/dto"
	"bingo-example/domain/entity/user"
)

// UserRep 用户响应
type UserRep struct{}

func (r *UserRep) User2Profile(user *user.User) *dto.Profile {
	return &dto.Profile{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Phone:     user.Phone,
		Email:     user.Email,
		Birthday:  user.Profile.Birth(),
		Gender:    user.Profile.GenderString(),
		Level:     user.Profile.Level,
		Signature: user.Profile.Signature,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (r *UserRep) User2ProfileMap(user *user.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"phone":      user.Phone,
		"email":      user.Email,
		"nickname":   user.Nickname,
		"avatar":     user.Avatar,
		"birthday":   user.Profile.Birth(),
		"gender":     user.Profile.GenderString(),
		"level":      user.Profile.Level,
		"signature":  user.Profile.Signature,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
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
