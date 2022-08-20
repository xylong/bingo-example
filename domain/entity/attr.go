package entity

type (
	// Attr 模型属性设置方法
	Attr func(interface{})

	// Attrs 属性方法集合
	Attrs []Attr
)

// Apply 为模型设置属性
func (a Attrs) Apply(model interface{}) {
	for _, attr := range a {
		attr(model)
	}
}
