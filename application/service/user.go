package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/domain/aggregate"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	Jwt *JwtService `inject:"-"`

	Req *assembler.UserReq
	Rep *assembler.UserRep

	DB *gorm.DB `inject:"-"`
}

func (s *UserService) Index() string {
	return "aa"
}

// Register 注册
func (s *UserService) Register(param *dto.RegisterParam) (int, string, string) {
	tx := s.DB.Begin()

	member := new(aggregate.Member).Builder(s.Req.Register2User(param)).
		SetProfile(s.Req.Register2Profile(param)).SetUserRepo(tx).SetProfileRepo(tx).Build()
	if err := member.Create(); err != nil {
		tx.Rollback()
		return 1001, err.Error(), ""
	} else {
		tx.Commit()
	}

	token, err := s.Jwt.generateToken(member.User.ID)
	if err != nil {
		return 1001, err.Error(), ""
	}

	return 0, "", token
}

func (s *UserService) Login(param *dto.LoginParam) (int, string, string) {
	member := new(aggregate.Member).Builder(s.Req.Login2User(param)).SetUserRepo(s.DB).Build()
	if err := member.Get("Profile"); err != nil {
		return 1002, err.Error(), ""
	}

	if !member.User.Profile.VerifyPassword(param.Password) {
		return 1002, "账号或密码错误", ""
	}

	token, err := s.Jwt.generateToken(member.User.ID)
	if err != nil {
		return 1001, err.Error(), ""
	}

	return 0, "", token
}
