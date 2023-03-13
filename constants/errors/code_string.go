// Code generated by "stringer -type Code -linecomment"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-0]
	_ = x[ParamError-400]
	_ = x[Unauthorized-1001]
	_ = x[PasswordError-1002]
	_ = x[RegisterError-1003]
}

const (
	_Code_name_0 = "OK"
	_Code_name_1 = "参数错误"
	_Code_name_2 = "未授权帐号或密码错误注册失败"
)

var (
	_Code_index_2 = [...]uint8{0, 9, 30, 42}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case i == 400:
		return _Code_name_1
	case 1001 <= i && i <= 1003:
		i -= 1001
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}