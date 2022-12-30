package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/constants"
	"bingo-example/domain/entity/book"
	"bingo-example/infrastructure/dao"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/olivere/elastic/v7"
	"github.com/xylong/bingo/ioc"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

// BookService 📚
type BookService struct {
	Req *assembler.BookReq `inject:"-"`
	Rep *assembler.BookRep `inject:"-"`

	DB *gorm.DB        `inject:"-"`
	Es *elastic.Client `inject:"-"`
}

// BatchImport 批量导入
func (s *BookService) BatchImport(ctx context.Context) {
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
				req.Index(constants.BookIndex).Id(strconv.Itoa(b.ID)).Doc(b)
				bulk.Add(req)
			}

			if rep, err := bulk.Do(ctx); err != nil {
				zap.L().Error("import book failed", zap.Error(err))
			} else {
				zap.L().Info("import book succeed", zap.Any("books", rep))
			}
		}()

		page++
	}

	wg.Wait()
}

// Search 搜索
func (s *BookService) Search(ctx context.Context, param *dto.BookSearchParam) *dto.CountList {
	// 指定字段
	var fields = []string{"id", "name", "author", "press", "date", "price1", "price2"}

	result, err := s.Es.Search().Index(constants.BookIndex).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include(fields...)).
		Query(s.Req.Filter(param)).SortBy(s.Req.Sort(param.Sorts)...).
		From(param.Offset()).Size(param.PageSize).
		Do(ctx)

	if err != nil {
		zap.L().Error("search book error", zap.Error(err))
		return nil
	}

	return s.Rep.EsSearchResult2CountList(result)
}

// GetPress 获取出版社
func (s *BookService) GetPress(ctx context.Context) []interface{} {
	collapse := elastic.NewCollapseBuilder(constants.BookPress)
	res, err := s.Es.Search().Index(constants.BookIndex).
		Collapse(collapse).FetchSource(false).Size(20).
		Do(ctx)

	if err != nil {
		zap.L().Error("get book press error", zap.Error(err))
	}

	return s.Rep.Fields2Slice(res, constants.BookPress)
}

func (s *BookService) GraphSearch(ctx context.Context) interface{} {
	param := graphql.Params{
		Schema:        s.GraphSchema(ctx),
		RequestString: constants.BookRequest,
	}

	result := graphql.Do(param)
	if result.HasErrors() {
		zap.L().Error("graph search", zap.Any("param", param), zap.Any("error", result.Errors))
		return nil
	}

	return result
}

func (s *BookService) GraphSchema(ctx context.Context) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: s.graphQuery(ctx),
	})

	if err != nil {
		zap.L().Error("book schema", zap.Error(err))
	}

	return schema
}

func (s *BookService) graphQuery(ctx context.Context) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BookQuery",
		Fields: graphql.Fields{
			"Book": &graphql.Field{
				Type: book.Graph(),
				Args: map[string]*graphql.ArgumentConfig{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if v, ok := p.Args["id"]; ok {
						db := ioc.Factory.Get((*gorm.DB)(nil))
						return dao.NewBookRepo(db.(*gorm.DB)).GetByID(v.(int))
					} else {
						return nil, nil
					}
				},
			},
			"Search": &graphql.Field{
				Type: graphql.NewList(book.Graph()),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if result, err := s.Es.Search().Index(constants.BookIndex).
						Query(s.Req.WildcardName(p.Args["name"].(string))).
						Do(ctx); err != nil {
						return nil, err
					} else {
						return s.Rep.Result2Books(result), nil
					}
				},
			},
		},
	})
}

// GetByID 详情
func (s *BookService) GetByID(ctx context.Context, id string) interface{} {
	if res, err := s.Es.Get().Index(constants.BookIndex).Id(id).Do(ctx); err != nil {
		zap.L().Warn("not found", zap.Error(err), zap.String("id", id))
		return nil
	} else {
		return res.Source
	}
}

// Create 创建
func (s *BookService) Create(ctx context.Context, param *dto.BookStoreParam) error {
	b := s.Req.StoreParam2Book(param)
	if err := dao.NewBookRepo(s.DB).Create(b); err != nil {
		zap.L().Error("create book", zap.Error(err), zap.Any("book", b))
		return fmt.Errorf("创建失败")
	}

	if _, err := s.Es.Index().Index(constants.BookIndex).
		Id(strconv.Itoa(b.ID)).BodyJson(b).
		Do(ctx); err != nil {
		zap.L().Error("create book", zap.Error(err))
	}

	return nil
}

// Update 更新
func (s *BookService) Update(ctx context.Context, request *dto.BookUrlRequest, param *dto.BookStoreParam) error {
	b := s.Req.StoreParam2Book(param, request)
	if err := dao.NewBookRepo(s.DB).Update(b); err != nil {
		zap.L().Error("update book", zap.Error(err), zap.Any("book", b))
		return fmt.Errorf("更新失败")
	}

	if _, err := s.Es.Update().Index(constants.BookIndex).
		Id(strconv.Itoa(b.ID)).Doc(b).Refresh("true").
		Do(ctx); err != nil {
		zap.L().Error("update book", zap.Error(err))
	}

	return nil
}

// Delete 删除
func (s *BookService) Delete(ctx context.Context, request *dto.BookUrlRequest) error {
	if err := dao.NewBookRepo(s.DB).Delete(request.ID); err != nil {
		zap.L().Error("delete book", zap.Error(err), zap.Any("id", request.ID))
		return fmt.Errorf("删除失败")
	}

	if _, err := s.Es.Delete().Index(constants.BookIndex).Id(strconv.Itoa(request.ID)).Refresh("true").Do(ctx); err != nil {
		zap.L().Error("delete book", zap.Error(err), zap.Any("id", request.ID))
		return fmt.Errorf("删除失败")
	}

	return nil
}
