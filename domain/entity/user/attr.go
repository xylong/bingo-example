package user

type (
	// Attr 模型属性设置方法
	Attr func(*User)

	// Attrs 属性方法集合
	Attrs []Attr
)

// Apply 为模型设置属性
func (a Attrs) Apply(user *User) {
	for _, attr := range a {
		attr(user)
	}
}

func WithID(id int) Attr {
	return func(user *User) {
		if id > 0 {
			user.ID = id
		}
	}
}

func WithPhone(phone string) Attr {
	return func(user *User) {
		if phone != "" {
			user.Phone = phone
		}
	}
}
