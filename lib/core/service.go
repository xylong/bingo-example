package core

import (
	"bingo-example/application/service"
	. "bingo-example/lib/factory"
)

// Service 服务
type Service struct{}

func NewService() *Service {
	return &Service{}
}

// User 创建UserService
func (s *Service) User() *service.UserService {
	return new(ServiceFactory).Create(User).(*service.UserService)
}
