package dao

import (
	"bingo-example/domain/entity/profile"
	"bingo-example/domain/repository"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfileRepo(db *gorm.DB) repository.IProfileRepo {
	return &ProfileRepo{db: db}
}

func (r *ProfileRepo) Create(profile *profile.Profile) error {
	return r.db.Create(profile).Error
}
