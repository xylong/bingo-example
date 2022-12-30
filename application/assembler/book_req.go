package assembler

import (
	"bingo-example/application/dto"
	. "bingo-example/constants"
	"bingo-example/domain/entity/book"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

const (
	scoreDesc = iota // 匹配度降序
	priceAsc         // 价格升序
	priceDesc        // 价格降序
	dateAsc          // 日期升序
	dateDesc         // 日期降序
)

var (
	// 排序规则
	sortSchema = map[uint8]*elastic.FieldSort{
		BookPrice1Desc: elastic.NewFieldSort(BookPrice1).Desc(),
		BookPrice1Asc:  elastic.NewFieldSort(BookPrice1).Asc(),
		BookPrice2Desc: elastic.NewFieldSort(BookPrice2).Desc(),
		BookPrice2Asc:  elastic.NewFieldSort(BookPrice2).Asc(),
		BookDateDesc:   elastic.NewFieldSort(BookDate).Desc(),
		BookDateAsc:    elastic.NewFieldSort(BookDate).Asc(),
	}
)

type BookReq struct {
	*ApiRequest
}

// Filter 过滤
func (r *BookReq) Filter(param *dto.BookSearchParam) *elastic.BoolQuery {
	var queries []elastic.Query

	{
		if param.Name != "" {
			queries = append(queries, r.NameQuery(param.Name))
		}

		if param.Press != "" {
			queries = append(queries, r.PressQuery(param.Press))
		}

		if param.Lowest > 0 || param.Highest > 0 {
			queries = append(queries, r.PriceQuery(param.Lowest, param.Highest))
		}
	}

	return elastic.NewBoolQuery().Must(queries...)
}

// NameQuery 书名检索
func (r BookReq) NameQuery(name string) *elastic.MatchQuery {
	return elastic.NewMatchQuery(BookName, name)
}

// WildcardName 匹配书名
func (r *BookReq) WildcardName(pattern string) *elastic.WildcardQuery {
	return elastic.NewWildcardQuery(BookName, pattern)
}

// PressQuery 出版社检索
func (r *BookReq) PressQuery(press string) *elastic.TermsQuery {
	var (
		arr []string
		brr []interface{}
	)

	arr = strings.Split(press, ",")
	if len(arr) > 0 {
		for _, s := range arr {
			brr = append(brr, s)
		}
	}

	return elastic.NewTermsQuery(BookPress, brr...)
}

// PriceQuery 价格区间检索
func (r *BookReq) PriceQuery(lowPrice, highPrice float64) *elastic.RangeQuery {
	query := elastic.NewRangeQuery(BookPrice1)

	if lowPrice > 0 {
		query.Gte(lowPrice)
	}

	if highPrice > 0 {
		query.Lte(highPrice)
	}

	return query
}

// Sort 排序
func (r *BookReq) Sort(sort string) []elastic.Sorter {
	var sorts []elastic.Sorter

	arr := strings.Split(sort, ",")
	if len(arr) > 0 {
		for _, s := range arr {
			i, _ := strconv.Atoi(s)
			if v, ok := sortSchema[uint8(i)]; ok {
				sorts = append(sorts, v)
			}
		}
	} else {
		sorts = append(sorts, elastic.NewScoreSort().Desc())
	}

	return sorts
}

// StoreParam2Book 书籍表单参数转模型
func (r *BookReq) StoreParam2Book(param *dto.BookStoreParam, request ...*dto.BookUrlRequest) *book.Book {
	attrs := []book.Attribute{
		book.WithName(param.Name),
		book.WithBlurb(param.Blurb),
		book.WithPrice1(param.Price1),
		book.WithPrice2(param.Price2),
		book.WithAuthor(param.Author),
		book.WithPress(param.Press),
		book.WithDate(param.Press),
		book.WithKind(param.Kind),
	}

	if len(request) > 0 {
		attrs = append(attrs, book.WithID(request[0].ID))
	}

	return book.New(attrs...)
}
