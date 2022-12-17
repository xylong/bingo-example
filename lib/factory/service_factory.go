package factory

import (
	"bingo-example/application/assembler"
	"bingo-example/application/service"
)

const (
	Jwt CreateType = iota // 用户
	User
	Book
	Question
)

// ServiceFactory 服务工厂
type ServiceFactory struct{}

// Create 创建service
func (*ServiceFactory) Create(createType CreateType) interface{} {
	switch createType {
	case Jwt:
		return &service.JwtService{}
	case User:
		return &service.UserService{Req: &assembler.UserReq{}, Rep: &assembler.UserRep{}}
	case Book:
		return &service.BookService{Req: &assembler.BookReq{}, Rep: &assembler.BookRep{}}
	case Question:
		return &service.QuestionService{Req: &assembler.QuestionReq{}, Rep: assembler.QuestionRep{}}
	default:
		return nil
	}
}
