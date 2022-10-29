package core

import (
	"bingo-example/application/service"
	"bingo-example/application/service/common"
	. "bingo-example/lib/factory"
)

// Service 服务
type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Es 创建elastic search服务
func (s Service) Es() *common.ElasticSearch {
	return new(ServiceFactory).Create(Es).(*common.ElasticSearch)
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
