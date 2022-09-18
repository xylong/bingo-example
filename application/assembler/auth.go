package assembler

import "golang.org/x/crypto/bcrypt"

type AuthReq struct{}

// EncryptPassword 加密密码
func (r *AuthReq) EncryptPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pwd), err
}

type AuthRep struct {
}
