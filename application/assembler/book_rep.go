package assembler

import (
	"bingo-example/domain/entity"
	"github.com/olivere/elastic/v7"
	"reflect"
)

type BookRep struct{}

// Result2Slice elastic结果转为slice
func (r *BookRep) Result2Slice(result *elastic.SearchResult) []*entity.Book {
	var (
		book  *entity.Book
		books []*entity.Book
	)

	t := reflect.TypeOf(book)
	for _, b := range result.Each(t) {
		books = append(books, b.(*entity.Book))
	}

	return books
}
