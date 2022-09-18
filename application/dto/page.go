package dto

import (
	_ "bingo-example/application/validation"
)

// Pagination 分页
type Pagination struct {
	Page     int `json:"page" form:"page" binding:"required,gte=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,gte=1,lte=100"`
}
