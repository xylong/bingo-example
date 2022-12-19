package g

import (
	. "bingo-example/domain/entity/book"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) GetByID(id int) (*Book, error) {
	book := &Book{}
	err := r.db.First(book, id).Error
	return book, err
}

func (r *BookRepo) Get() ([]*Book, error) {
	var books []*Book
	err := r.db.Limit(10).Offset(1).Find(&books).Error
	return books, err
}
