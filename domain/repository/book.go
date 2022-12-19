package repository

import "bingo-example/domain/entity/book"

type IBookRepo interface {
	GetByID(int) (*book.Book, error)
}
