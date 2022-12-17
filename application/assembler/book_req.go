package assembler

import (
	"bingo-example/application/dto"
	. "bingo-example/constants"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
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

type BookReq struct{}

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
