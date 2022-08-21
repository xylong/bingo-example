package server

import "bingo-example/application/assembler"

// UserService 用户服务
type UserService struct {
	Req *assembler.UserReq
	Rep *assembler.UserRep
}

func (s *UserService) Index() string {
	return "aa"
}
