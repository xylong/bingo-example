package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/application/utils/token"
	"bingo-example/constants"
	"bingo-example/constants/errors"
	"bingo-example/domain/aggregate"
	"bingo-example/domain/entity/user"
	"bingo-example/infrastructure/dao"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// UserService 用户服务
type UserService struct {
	Req *assembler.UserReq
	Rep *assembler.UserRep

	DB  *gorm.DB      `inject:"-"`
	Rdb *redis.Client `inject:"-"`
}

func (s *UserService) Index() string {
	return "aa"
}

// Register 注册
func (s *UserService) Register(param *dto.RegisterParam) map[string]string {
	tx := s.DB.Begin()

	member := new(aggregate.Member).Builder(s.Req.Register2User(param)).
		SetProfile(s.Req.Register2Profile(param)).SetUserRepo(tx).SetProfileRepo(tx).Build()
	if err := member.Create(); err != nil {
		tx.Rollback()
		return nil
	} else {
		tx.Commit()
	}

	accessToken, refreshToken, err := token.Generate(member.User.ID)
	if err != nil {
		return nil
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
}

// Login 登录
func (s *UserService) Login(param *dto.LoginParam) (int, string, map[string]string) {
	member := new(aggregate.Member).Builder(s.Req.Login2User(param)).SetUserRepo(s.DB).Build()
	if err := member.Take(map[string][]string{"": []string{"id"}, "Profile": []string{"user_id", "password", "salt"}}); err != nil {
		return errors.PasswordError.Int(), errors.PasswordError.String(), nil
	}

	if !member.User.Profile.VerifyPassword(param.Password) {
		return errors.PasswordError.Int(), errors.PasswordError.String(), nil
	}

	accessToken, refreshToken, err := token.Generate(member.User.ID)
	if err != nil {
		return errors.Unauthorized.Int(), errors.Unauthorized.String(), nil
	}

	return 0, "", map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
}

// Profile 个人信息
func (s *UserService) Profile(id int) (*dto.Profile, error) {
	// 1.从缓存取信息
	profile := &dto.Profile{}
	ctx := context.Background()
	key := fmt.Sprintf(constants.ProfileCache, id)

	if err := s.Rdb.HGetAll(ctx, key).Scan(profile); err != nil {
		zap.L().Error("scan profile", zap.Error(err))
	}

	// 2.缓存没有再从数据库取，取完设置缓存
	if profile.ID == 0 {
		u := user.New(user.WithID(id))
		if err := new(aggregate.Member).Builder(u).SetUserRepo(s.DB).Build().Take(map[string][]string{
			"":        []string{"id", "phone", "email", "nickname", "avatar", "created_at"},
			"Profile": []string{"user_id", "birthday", "gender", "level", "signature"}}); err != nil {
			return nil, err
		}

		profile = s.Rep.User2Profile(u)
		if ok, err := s.Rdb.HMSet(ctx, key, s.Rep.User2ProfileMap(u)).Result(); err != nil || !ok {
			zap.L().Error("cache profile", zap.Error(err))
		}

		if !s.Rdb.Expire(ctx, key, time.Minute).Val() {
			zap.L().Warn("profile cache expire")
		}
	}

	return profile, nil
}

// Get 获取用户
func (s *UserService) Get(ctx context.Context, request *dto.UserRequest) (*dto.SimpleUserList, error) {
	total, users, err := new(aggregate.Member).Builder(user.New()).SetUserRepo(s.DB).Build().Get(request)
	if err != nil {
		return nil, err
	}

	return &dto.SimpleUserList{
		Total: total,
		List:  s.Rep.SimpleList(users),
	}, nil
}

// CountReg 按月统计注册
func (s *UserService) CountReg(ctx context.Context, request *dto.RegisterCountRequest) interface{} {
	result, err := dao.NewUserRepo(s.DB).CountRegister([]*dto.RegisterCount{}, request.Month)
	if err != nil {
		zap.L().Error("count", zap.Error(err))
	}

	for _, item := range result.([]*dto.RegisterCount) {
		item.Date = item.Date[:10]
	}

	return result
}
