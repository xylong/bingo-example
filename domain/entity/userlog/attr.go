package userlog

type (
	Attr  func(*UserLog)
	Attrs []Attr
)

func (a Attrs) Apply(log *UserLog) {
	for _, attr := range a {
		attr(log)
	}
}
