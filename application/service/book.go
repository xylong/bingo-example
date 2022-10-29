package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/domain/entity"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

const (
	bookIndex = "books" // elastic索引
)

type BookService struct {
	Req *assembler.BookReq `inject:"-"`
	Rep *assembler.BookRep `inject:"-"`

	DB *gorm.DB        `inject:"-"`
	Es *elastic.Client `inject:"-"`
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
func (s *BookService) Search(query *dto.BookQuery) interface{} {
	var (
		result *elastic.SearchResult
		err    error
	)

	if query.Press != "" {
		term := elastic.NewTermsQuery("book_press", s.Req.FilterPress(query.Press)...)
		result, err = s.Es.Search().Query(term).Index(bookIndex).Do(context.Background())
	} else {
		result, err = s.Es.Search().Index(bookIndex).Do(context.Background())
	}

	if err != nil {
		zap.L().Error("search book error", zap.Error(err))
		return nil
	}

	return s.Rep.Result2Slice(result)
}
