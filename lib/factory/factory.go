package factory

// CreateType 创建类型
type CreateType uint8

// Creator 总工厂
type Creator interface {
	Create(CreateType) interface{}
}
