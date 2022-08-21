package configuration

import (
	"bingo-example/application/server"
	"bingo-example/lib/factory"
)

// Service 服务
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// User 创建UserService
func (s *Service) User() *server.UserService {
	return new(factory.ServiceFactory).Create(factory.User).(*server.UserService)
}
