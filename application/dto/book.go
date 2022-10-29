package dto

type (
	// BookSearchParam 📚查询参数
	BookSearchParam struct {
		*Pagination
		Name     string  `form:"name" json:"name" binding:"omitempty"`                                       // 书名
		Press    string  `form:"press" json:"press" binding:"omitempty"`                                     // 出版社
		Lowest   float64 `form:"lowest" json:"lowest" binding:"omitempty,gte=0,lte=10000"`                   // 最低价
		Highest  float64 `form:"highest" json:"highest" binding:"omitempty,gte=0,lte=10000,gtefield=Lowest"` // 最高价
		OrderSet struct {
			Score bool `form:"score" json:"score"`
			Price int  `form:"price" json:"price" binding:"oneof=0 1 2"`
		} `json:"order_set" binding:"required,dive"`
	}
)
