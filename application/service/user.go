package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/constants"
	"bingo-example/domain/aggregate"
	"bingo-example/domain/entity/user"
	"bingo-example/infrastructure/dao"
	"bingo-example/utils"
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
func (s *UserService) Register(param *dto.RegisterParam) (int, string, string) {
	tx := s.DB.Begin()

	member := new(aggregate.Member).Builder(s.Req.Register2User(param)).
		SetProfile(s.Req.Register2Profile(param)).SetUserRepo(tx).SetProfileRepo(tx).Build()
	if err := member.Create(); err != nil {
		tx.Rollback()
		return 1001, err.Error(), ""
	} else {
		tx.Commit()
	}

	token, err := utils.GenerateToken(member.User.ID)
	if err != nil {
		return 1001, err.Error(), ""
	}

	return 0, "", token
}

// Login 登录
func (s *UserService) Login(param *dto.LoginParam) (int, string, string) {
	member := new(aggregate.Member).Builder(s.Req.Login2User(param)).SetUserRepo(s.DB).Build()
	if err := member.Take(map[string][]string{"Profile": []string{"user_id", "password", "salt"}}); err != nil {
		return 1002, err.Error(), ""
	}

	if !member.User.Profile.VerifyPassword(param.Password) {
		return 1002, "账号或密码错误", ""
	}

	token, err := utils.GenerateToken(member.User.ID)
	if err != nil {
		return 1001, err.Error(), ""
	}

	return 0, "", token
}

// Profile 个人信息
func (s *UserService) Profile(id int) (int, string, *dto.Profile) {
	profile := &dto.Profile{}
	ctx := context.Background()
	key := fmt.Sprintf(constants.ProfileCache, id)

	if err := s.Rdb.HGetAll(ctx, key).Scan(profile); err != nil {
		zap.L().Error("scan profile", zap.Error(err))
	}

	if profile.ID == 0 {
		u := user.New(user.WithID(id))
		if err := new(aggregate.Member).Builder(u).SetUserRepo(s.DB).Build().Take(map[string][]string{
			"":        []string{"id", "phone", "email", "nickname", "avatar", "created_at"},
			"Profile": []string{"user_id", "birthday", "gender", "level", "signature"}}); err != nil {
			return 1002, "not found", nil
		}

		profile = s.Rep.User2Profile(u)
		if ok, err := s.Rdb.HMSet(ctx, key, s.Rep.User2ProfileMap(u)).Result(); err != nil || !ok {
			zap.L().Error("cache profile", zap.Error(err))
		}

		if !s.Rdb.Expire(ctx, key, time.Minute).Val() {
			zap.L().Warn("profile cache expire")
		}
	}

	return 0, "", profile
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
