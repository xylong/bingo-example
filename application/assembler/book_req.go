package assembler

import (
	"bingo-example/application/dto"
	"strings"
)

type BookReq struct{}

// Filter 组合过滤条件
func (r *BookReq) Filter(query *dto.BookQuery) {
	if query.Press != "" {

	}
}

// FilterPress 根据出版社过滤
func (r *BookReq) FilterPress(press string) []interface{} {
	arr := strings.Split(press, ",")
	if len(arr) == 0 {
		return nil
	}

	var condition []interface{}
	for _, item := range arr {
		condition = append(condition, item)
	}

	return condition
}
