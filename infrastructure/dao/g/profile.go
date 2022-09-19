package g

import (
	"bingo-example/domain/entity/profile"
	"gorm.io/gorm"
)

type ProfileDao struct {
	db *gorm.DB
}

func NewProfileDao(db *gorm.DB) *ProfileDao {
	return &ProfileDao{db: db}
}

func (d *ProfileDao) Create(profile *profile.Profile) error {
	return d.db.Create(profile).Error
}
