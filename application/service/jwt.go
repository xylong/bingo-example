package service

import (
	"bingo-example/domain/entity"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtService struct {
}

//func (s *JwtService) Login(param *dto.LoginParam) string {
//	u := user.NewUser(user.WithPhone(param.Phone))
//	if err := g.NewUserRepo(s.DB).Get(u); err != nil {
//		return ""
//	}
//
//	if !u.Profile.VerifyPassword(param.Password) {
//		return ""
//	}
//
//	token, err := s.generateToken(u.ID)
//	if err != nil {
//		return ""
//	}
//
//	return token
//}

// generateToken 生成token
func (s *JwtService) generateToken(id int) (string, error) {
	claims := &entity.JwtClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(entity.TokenExpireDuration)),
			Issuer:    "bingo-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 用指定方法创建签名对象
	return token.SignedString(entity.JwtSecret)                // 使用指定的secret签名并获得完整的编码后的字符串token
}

// ParseToken 解析token
func (s *JwtService) ParseToken(token string) (*entity.JwtClaims, error) {
	if _token, err := jwt.ParseWithClaims(token, &entity.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return entity.JwtSecret, nil
	}); err != nil {
		return nil, err
	} else {
		// 对token对象中的Claim进行类型断言
		if claims, ok := _token.Claims.(*entity.JwtClaims); ok && _token.Valid {
			return claims, nil
		}
	}

	return nil, fmt.Errorf("invalid token")
}
