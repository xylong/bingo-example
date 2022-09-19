package factory

import (
	"bingo-example/application/assembler"
	"bingo-example/application/service"
)

const (
	Jwt CreateType = iota // 用户
	User
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
	default:
		return nil
	}
}
