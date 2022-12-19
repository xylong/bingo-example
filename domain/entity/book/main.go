package book

import (
	"github.com/graphql-go/graphql"
)

const (
	OrderByPriceAsc  = iota + 1 // 价格从低到高
	OrderByPriceDesc            // 价格从高到低
)

type Books []*Book

// Book 书籍
type Book struct {
	ID      int     `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Name    string  `gorm:"column:name;type:varchar(50);not null;" json:"name"`  // 名称
	Blurb   string  `gorm:"column:blurb;type:text" json:"blurb"`                 // 简介
	Price1  float64 `gorm:"column:price1;type:decimal;default:0;" json:"price1"` // 价格1
	Price2  float64 `gorm:"column:price2;type:decimal;default:0;" json:"price2"` // 价格2
	Author  string  `gorm:"column:author;type:varchar(50)" json:"author"`        // 作者
	Press   string  `gorm:"column:press;type:varchar(50)" json:"press"`          // 出版社
	Date    string  `gorm:"column:date;type:char(11)" json:"date"`               // 出版时间
	Kind    uint8   `gorm:"column:kind;type:int" json:"kind"`                    // 分类id
	KindStr string  `gorm:"column:kind_str;type:varchar(15)" json:"kind_str"`    // 分类名
}

func New(attributes ...Attribute) *Book {
	b := &Book{}
	Attributes(attributes).Apply(b)
	return b
}

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
