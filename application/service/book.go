package service

import (
	"bingo-example/application/assembler"
	"bingo-example/application/dto"
	"bingo-example/constants"
	"bingo-example/domain/entity/book"
	"bingo-example/infrastructure/dao/g"
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
				req.Index(constants.BookIndex).Id(strconv.Itoa(b.ID)).Doc(b)
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
	result, err := s.Es.Search().Index(constants.BookIndex).
		Query(s.Req.Filter(param)).SortBy(s.Req.Sort(param.Sorts)...).
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
	collapse := elastic.NewCollapseBuilder(constants.BookPress)
	res, err := s.Es.Search().Index(constants.BookIndex).
		Collapse(collapse).FetchSource(false).Size(20).
		Do(context.Background())

	if err != nil {
		zap.L().Error("get book press error", zap.Error(err))
	}

	return s.Rep.Fields2Slice(res, constants.BookPress)
}

func (s *BookService) GraphSearch() interface{} {
	param := graphql.Params{
		Schema:        s.GraphSchema(),
		RequestString: constants.BookRequest,
	}

	result := graphql.Do(param)
	if result.HasErrors() {
		zap.L().Error("graph search", zap.Any("param", param), zap.Any("error", result.Errors))
		return nil
	}

	return result
}

func (s *BookService) GraphSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: s.graphQuery(),
	})

	if err != nil {
		zap.L().Error("book schema", zap.Error(err))
	}

	return schema
}

func (s *BookService) graphQuery() *graphql.Object {
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
						return g.NewBookRepo(db.(*gorm.DB)).GetByID(v.(int))
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
						Do(context.Background()); err != nil {
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
func (s *BookService) GetByID(id string) interface{} {
	if res, err := s.Es.Get().Index(constants.BookIndex).Id(id).Do(context.Background()); err != nil {
		zap.L().Warn("not found", zap.Error(err), zap.String("id", id))
		return nil
	} else {
		return res.Source
	}
}

// Create 创建
func (s *BookService) Create(param *dto.BookStoreParam) error {
	b := s.Req.Param2Book(param)
	err := g.NewBookRepo(s.DB).Create(b)
	if err != nil {
		zap.L().Error("create book", zap.Error(err), zap.Any("book", b))
		return fmt.Errorf("创建失败")
	}

	if _, err = s.Es.Index().Index(constants.BookIndex).Id(strconv.Itoa(b.ID)).BodyJson(b).Do(context.Background()); err != nil {
		zap.L().Error("create book", zap.Error(err))
	}

	return nil
}
