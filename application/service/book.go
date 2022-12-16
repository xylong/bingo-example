package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/domain/entity/book"
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type BookService struct {
	Req *assembler.BookReq `inject:"-"`
	Rep *assembler.BookRep `inject:"-"`

	DB *gorm.DB        `inject:"-"`
	Es *elastic.Client `inject:"-"`
}

// BatchImport 批量导入
func (s *BookService) BatchImport() {
	page, pageSize := 1, 1000
	wg := sync.WaitGroup{}

	for {
		// 从mysql获取数据
		books := book.Books{}
		err := s.DB.Model(&book.Book{}).Order("id desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&books).Error
		if err != nil || len(books) == 0 {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			// 导入到es
			bulk := s.Es.Bulk()
			for _, b := range books {
				req := elastic.NewBulkIndexRequest()
				req.Index(assembler.BookIndex).Id(strconv.Itoa(b.ID)).Doc(b)
				bulk.Add(req)
			}

			if rep, err := bulk.Do(context.Background()); err != nil {
				zap.L().Error("import book failed", zap.Error(err))
			} else {
				zap.L().Info("import book succeed", zap.Any("books", rep))
			}
		}()

		page++
	}

	wg.Wait()
}

// Search 📚搜索
func (s *BookService) Search(param *dto.BookSearchParam) interface{} {
	result, err := s.Es.Search().Index(assembler.BookIndex).
		Query(s.Req.Query(param)).SortBy(s.Req.Sort(param.Sorts)...).
		From(param.Offset()).Size(param.PageSize).
		Do(context.Background())
	if err != nil {
		zap.L().Error("search book error", zap.Error(err))
		return nil
	}

	return s.Rep.Result2Slice(result)
}

// GetPress 获取出版社
func (s *BookService) GetPress() []interface{} {
	collapse := elastic.NewCollapseBuilder(assembler.BookPress)
	res, err := s.Es.Search().Index(assembler.BookIndex).Size(20).Collapse(collapse).FetchSource(false).Do(context.Background())
	if err != nil {
		zap.L().Error("get book press error", zap.Error(err))
	}

	return s.Rep.Fields2Slice(res, assembler.BookPress)
}
