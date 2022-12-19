package graph

import (
	"bingo-example/infrastructure/dao/g"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/xylong/bingo/ioc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Graph 实体映射图表
func Graph() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BookModel",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"name":     &graphql.Field{Type: graphql.String},
			"blurb":    &graphql.Field{Type: graphql.String},
			"price1":   &graphql.Field{Type: graphql.Float},
			"price2":   &graphql.Field{Type: graphql.Float},
			"author":   &graphql.Field{Type: graphql.String},
			"press":    &graphql.Field{Type: graphql.String},
			"date":     &graphql.Field{Type: graphql.String},
			"kind":     &graphql.Field{Type: graphql.Int},
			"kind_str": &graphql.Field{Type: graphql.String},
		},
		Description: "📚",
	})
}

// Query 创建图表查询对象
func Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BookQuery",
		Fields: graphql.Fields{
			"Book": &graphql.Field{
				Type: Graph(),
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
						return nil, errors.New("book id param error")
					}
				},
			},
		},
	})
}

// Schema 创建图表查询规则
func Schema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query(),
	})

	if err != nil {
		zap.L().Error("book schema", zap.Error(err))
	}

	return schema
}
