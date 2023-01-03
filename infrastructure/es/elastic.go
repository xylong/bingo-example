package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type docData struct {
	ctx   context.Context
	index string
	id    string
	doc   interface{}
}

type Elastic struct {
	Es *elastic.Client `inject:"-"`

	writer chan *docData
}

func NewElastic() *Elastic {
	return &Elastic{
		writer: make(chan *docData, 16),
	}
}

// Upsert 修改文档，不存在则插入
func (e *Elastic) Upsert(ctx context.Context, index, id string, doc interface{}) {
	e.writer <- &docData{
		ctx:   ctx,
		index: index,
		id:    id,
		doc:   doc,
	}
}

func (e *Elastic) write() {
	for data := range e.writer {
		_, err := e.Es.Update().Index(data.index).Id(data.id).Doc(data.doc).Upsert(data.doc).Refresh("true").Do(data.ctx)
		zap.L().Error("upsert", zap.Error(err))
	}
}
