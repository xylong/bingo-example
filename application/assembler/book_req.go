package assembler

import (
	"bingo-example/application/dto"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

const (
	BookIndex  = "books"  // 📚es索引
	BookName   = "name"   // 书名
	BookBlurb  = "blurb"  // 简介
	BookPress  = "press"  // 出版社
	BookPrice1 = "price1" // 价格1
)

const (
	scoreDesc = iota // 匹配度降序
	priceAsc         // 价格升序
	priceDesc        // 价格降序
	dateAsc          // 日期升序
	dateDesc         // 日期降序
)

var (
	sortFunc = map[uint8]*elastic.FieldSort{
		priceAsc:  elastic.NewFieldSort("price1").Asc(),
		priceDesc: elastic.NewFieldSort("price1").Desc(),
		dateAsc:   elastic.NewFieldSort("date").Asc(),
		dateDesc:  elastic.NewFieldSort("date").Desc(),
	}
)

type BookReq struct{}

// Query es查询条件
func (r *BookReq) Query(param *dto.BookSearchParam) *elastic.BoolQuery {
	queries := make([]elastic.Query, 0)

	if param.Name != "" {
		queries = append(queries, r.MatchName(param.Name))
	}

	if param.Press != "" {
		queries = append(queries, r.InPress(param.Press))
	}

	if param.Lowest > 0 || param.Highest > 0 {
		queries = append(queries, r.ComparePrice(BookPrice1, param.Lowest, param.Highest))
	}

	return elastic.NewBoolQuery().Must(queries...)
}

// MatchName 匹配书籍名称
func (r *BookReq) MatchName(name string) *elastic.MatchQuery {
	return elastic.NewMatchQuery(BookName, name)
}

// MatchBlurb 匹配简介
func (r *BookReq) MatchBlurb(blurb string) *elastic.MatchQuery {
	return elastic.NewMatchQuery(BookBlurb, blurb)
}

// InPress 出版社
func (r *BookReq) InPress(presses string) *elastic.TermsQuery {
	var s []interface{}

	arr := strings.Split(presses, ",")
	for _, item := range arr {
		s = append(s, item)
	}

	return elastic.NewTermsQuery(BookPress, s...)
}

// ComparePrice 比较价格
func (r *BookReq) ComparePrice(field string, price1, price2 float64) *elastic.RangeQuery {
	query := elastic.NewRangeQuery(field)

	if price1 > 0 {
		query.Gte(price1)
	}

	if price2 > 0 {
		query.Lte(price2)
	}

	return query
}

// Sort 排序
func (r *BookReq) Sort(fields string) []elastic.Sorter {
	var sorts []elastic.Sorter

	keys := strings.Split(fields, ",")
	if len(keys) > 0 {
		for _, key := range keys {
			i, _ := strconv.Atoi(key)
			if v, ok := sortFunc[uint8(i)]; ok {
				sorts = append(sorts, v)
			} else {
				sorts = append(sorts, elastic.NewScoreSort().Desc())
			}
		}
	} else {
		sorts = append(sorts, elastic.NewScoreSort().Desc())
	}

	return sorts
}
