package initializers

import (
	"bingo-example/domain/entity/book"
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/entity/user"
	"bingo-example/domain/entity/userlog"
	"github.com/xylong/bingo/ioc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SyncDatabase 同步数据库
func SyncDatabase() {
	err := ioc.Factory.Get((*gorm.DB)(nil)).(*gorm.DB).AutoMigrate(
		&user.User{}, &profile.Profile{}, &userlog.UserLog{}, &book.Book{})

	if err != nil {
		zap.L().Warn("sync database", zap.Error(err))
	}
}
