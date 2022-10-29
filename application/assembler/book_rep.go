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

// Fields2Slice 将结果中指定字段转为slice
func (r *BookRep) Fields2Slice(result *elastic.SearchResult, key string) []interface{} {
	var res []interface{}

	for _, hit := range result.Hits.Hits {
		if v, ok := hit.Fields[key].([]interface{}); ok {
			res = append(res, v[0])
		} else {
			break
		}
	}

	return res
}
