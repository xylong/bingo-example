package common

import "github.com/olivere/elastic/v7"

// EsQuery elastic查询条件方法
type EsQuery []func(*elastic.Query) *elastic.SearchResult

// ElasticSearch es
type ElasticSearch struct {
	Es *elastic.Client `inject:"-"`
}

func (s *ElasticSearch) Must(query ...elastic.Query) *elastic.BoolQuery {
	return elastic.NewBoolQuery().Must(query...)
}
