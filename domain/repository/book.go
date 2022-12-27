package repository

import "bingo-example/domain/entity/book"

type IBookRepo interface {
	GetByID(int) (*book.Book, error)
	Create(*book.Book) error
	Update(*book.Book) error
	Delete(int) error
}
