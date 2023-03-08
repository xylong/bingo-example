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

// User 用户
func (s *Service) User() *service.UserService {
	return new(ServiceFactory).Create(User).(*service.UserService)
}

// Book 📚
func (s *Service) Book() *service.BookService {
	return new(ServiceFactory).Create(Book).(*service.BookService)
}

// Fruit 🍉
func (s *Service) Fruit() *service.FruitService {
	return new(ServiceFactory).Create(Fruit).(*service.FruitService)
}

// Question 题库
func (s *Service) Question() *service.QuestionService {
	return new(ServiceFactory).Create(Question).(*service.QuestionService)
}
