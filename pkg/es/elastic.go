package es

import (
	"github.com/olivere/elastic/v7"
)

var es *elastic.Client

// ES 获取elastic
func ES() *elastic.Client {
	return es
}

// Connect 连接es
func Connect(optionFunc ...elastic.ClientOptionFunc) {
	var err error

	if es, err = elastic.NewClient(optionFunc...); err != nil {
		panic(err)
	}
}
