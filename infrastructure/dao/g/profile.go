package g

import (
	"bingo-example/domain/entity/profile"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfileRepo(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (d *ProfileRepo) Create(profile *profile.Profile) error {
	return d.db.Create(profile).Error
}
