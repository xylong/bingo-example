package assembler

import "github.com/olivere/elastic/v7"

// ApiRequest api请求
type ApiRequest struct {
}

// SpecifyFields 指定es字段
func (r *ApiRequest) SpecifyFields(fields ...string) *elastic.FetchSourceContext {
	return elastic.NewFetchSourceContext(true).Include(fields...)
}
