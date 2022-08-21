package server

import (
	"bingo-example/application/assembler"
	"bingo-example/domain/aggregate"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
)

// UserService 用户服务
type UserService struct {
	Req *assembler.UserReq
	Rep *assembler.UserRep
}

func (s *UserService) Index() string {
	return "aa"
}

func (s *UserService) Create() {
	new(aggregate.Member).Builder(&user.User{}).SetProfile(&profile.Profile{}).Build()
}
