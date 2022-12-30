package assembler

import (
	"bingo-example/application/dto"
	"github.com/olivere/elastic/v7"
)

// ApiResponse api响应
type ApiResponse struct {
}

// EsSearchResult2CountList es搜索转带统计列表
func (r *ApiResponse) EsSearchResult2CountList(result *elastic.SearchResult) *dto.CountList {
	data := &dto.CountList{
		Total: result.TotalHits(),
		List:  nil,
	}

	if data.Total > 0 {
		for _, hit := range result.Hits.Hits {
			data.List = append(data.List, hit.Source)
		}
	}

	return data
}
