package profile

type (
	// Attr 模型属性设置方法
	Attr func(profile *Profile)

	// Attrs 属性方法集合
	Attrs []Attr
)

func (a Attrs) Apply(profile *Profile) {
	for _, attr := range a {
		attr(profile)
	}
}

func WithPassword(password string) Attr {
	return func(profile *Profile) {
		if password != "" {
			profile.Password = password
		}
	}
}
