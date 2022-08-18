package lib

import (
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/ioc"
	"sync"
)

// ConfigPool 配置
var ConfigPool = sync.Pool{
	New: func() interface{} {
		return ioc.Factory.Get((*bingo.Config)(nil))
	},
}
