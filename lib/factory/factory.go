package factory

// Type 创建类型
type Type int

// Creator 创建者
// 总工厂
type Creator interface {
	Create(Type) interface{}
}
