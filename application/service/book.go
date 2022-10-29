package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/application/service/common"
	"bingo-example/domain/entity"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

// es
const (
	bookIndex  = "books"       // 📚es索引
	bookName   = "book_name"   // 书名
	bookPress  = "book_press"  // 出版社
	bookPrice1 = "book_price1" // 价格1
)

type BookService struct {
	Req *assembler.BookReq `inject:"-"`
	Rep *assembler.BookRep `inject:"-"`

	DB *gorm.DB `inject:"-"`

	*common.ElasticSearch `inject:"-"`
}

// BatchImport 批量导入
func (s *BookService) BatchImport() {
	page, pageSize := 1, 500
	wg := sync.WaitGroup{}

	for {
		// 从mysql获取数据
		books := entity.Books{}
		err := s.DB.Model(&entity.Book{}).Order("book_id desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&books).Error
		if err != nil || len(books) == 0 {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			// 导入到es
			bulk := s.Es.Bulk()
			for _, book := range books {
				req := elastic.NewBulkIndexRequest()
				req.Index(bookIndex).Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(req)
			}
			rep, err := bulk.Do(context.Background())
			fmt.Println(rep, err)
		}()

		page++
	}

	wg.Wait()
}

// Search 📚搜索
func (s *BookService) Search(param *dto.BookSearchParam) interface{} {
	var (
		result  *elastic.SearchResult
		sort    []elastic.Sorter
		queries []elastic.Query
		err     error
	)

	// 过滤
	{
		if param.Name != "" {
			matchQuery := elastic.NewMatchQuery(bookName, param.Name)
			queries = append(queries, matchQuery)
		}

		if param.Press != "" {
			termsQuery := elastic.NewTermsQuery(bookPress, s.Req.FilterPress(param.Press)...)
			queries = append(queries, termsQuery)
		}

		if param.Lowest > 0 || param.Highest > 0 {
			rangeQuery := elastic.NewRangeQuery(bookPrice1)

			if param.Lowest > 0 {
				rangeQuery.Gte(param.Lowest)
			}

			if param.Highest > 0 {
				rangeQuery.Lte(param.Highest)
			}

			queries = append(queries, rangeQuery)
		}
	}

	// 排序
	{
		if param.OrderSet.Score {
			sort = append(sort, elastic.NewScoreSort().Desc())
		}
		if param.OrderSet.Price == entity.OrderByPriceAsc {
			sort = append(sort, elastic.NewFieldSort(bookPrice1).Asc())
		}
		if param.OrderSet.Price == entity.OrderByPriceDesc {
			sort = append(sort, elastic.NewFieldSort(bookPrice1).Desc())
		}
	}

	result, err = s.Es.Search().Index(bookIndex).Query(s.Must(queries...)).
		SortBy(sort...).From(param.Offset()).Size(param.PageSize).
		Do(context.Background())
	if err != nil {
		zap.L().Error("search book error", zap.Error(err))
		return nil
	}

	return s.Rep.Result2Slice(result)
}

// GetPress 获取出版社
func (s *BookService) GetPress() []interface{} {
	collapse := elastic.NewCollapseBuilder(bookPress)
	res, err := s.Es.Search().Index(bookIndex).Size(20).Collapse(collapse).FetchSource(false).Do(context.Background())
	if err != nil {
		zap.L().Error("get book press error", zap.Error(err))
	}

	return s.Rep.Fields2Slice(res, bookPress)
}
