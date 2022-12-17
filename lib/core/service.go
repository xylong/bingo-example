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

// Jwt 创建jwt
func (s *Service) Jwt() *service.JwtService {
	return new(ServiceFactory).Create(Jwt).(*service.JwtService)
}

// User 创建UserService
func (s *Service) User() *service.UserService {
	return new(ServiceFactory).Create(User).(*service.UserService)
}

// Book 书
func (s *Service) Book() *service.BookService {
	return new(ServiceFactory).Create(Book).(*service.BookService)
}

// Question 题库
func (s *Service) Question() *service.QuestionService {
	return new(ServiceFactory).Create(Question).(*service.QuestionService)
}
