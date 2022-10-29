package factory

import (
	"bingo-example/application/assembler"
	"bingo-example/application/service"
	"bingo-example/application/service/common"
)

const (
	Jwt CreateType = iota // 用户
	Es                    // elastic
	User
	Book
)

// ServiceFactory 服务工厂
type ServiceFactory struct{}

// Create 创建service
func (*ServiceFactory) Create(createType CreateType) interface{} {
	switch createType {
	case Jwt:
		return &service.JwtService{}
	case Es:
		return &common.ElasticSearch{}
	case User:
		return &service.UserService{Req: &assembler.UserReq{}, Rep: &assembler.UserRep{}}
	case Book:
		return &service.BookService{Req: &assembler.BookReq{}, Rep: &assembler.BookRep{}}
	default:
		return nil
	}
}
