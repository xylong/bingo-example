package common

import "github.com/olivere/elastic/v7"

type queries []func(*elastic.Query) *elastic.SearchResult

// ElasticSearch es
type ElasticSearch struct {
	Es *elastic.Client `inject:"-"`
}
