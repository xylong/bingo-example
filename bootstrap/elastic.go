package bootstrap

import (
	"bingo-example/pkg/config"
	"bingo-example/pkg/es"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// SetupElastic 设置elastic
func SetupElastic() {
	es.Connect(
		elastic.SetSniff(config.GetBool("elastic.sniff")),
		elastic.SetURL(fmt.Sprintf("http://%s:%v",
			config.Get("elastic.host"), config.Get("elastic.port"))))
}
