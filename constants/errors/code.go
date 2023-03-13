package errors

//go:generate stringer -type Code -linecomment
type Code uint16

func (i Code) Int() int {
	return int(i)
}

const (
	OK         Code = 0
	ParamError Code = 400 // 参数错误
	NotFound   Code = 404 // 未找到

	Unauthorized  Code = 1001 // 未授权
	PasswordError Code = 1002 // 帐号或密码错误
	RegisterError Code = 1003 // 注册失败
)
