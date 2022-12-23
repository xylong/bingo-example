package dto

type (
	// BookSearchParam 📚查询参数
	BookSearchParam struct {
		*Pagination
		Name    string  `form:"name" json:"name" binding:"omitempty"`                                       // 书名
		Press   string  `form:"press" json:"press" binding:"omitempty"`                                     // 出版社
		Lowest  float64 `form:"lowest" json:"lowest" binding:"omitempty,gte=0,lte=10000"`                   // 最低价
		Highest float64 `form:"highest" json:"highest" binding:"omitempty,gte=0,lte=10000,gtefield=Lowest"` // 最高价
		Sorts   string  `form:"sorts" json:"sorts" binding:"omitempty"`                                     // 排序
	}

	// BookStoreParam 📚创建参数
	BookStoreParam struct {
		Name   string  `form:"name" json:"name" binding:"required"`
		Blurb  string  `form:"blurb" json:"blurb" binding:"omitempty"` // 简介
		Price1 float64 `form:"price1" json:"price1" binding:"omitempty,gte=0"`
		Price2 float64 `form:"price2" json:"price2" binding:"omitempty,gte=0"`
		Author string  `form:"author" json:"author" binding:"omitempty"`                                                  // 作者
		Press  string  `form:"press" json:"press" binding:"omitempty"`                                                    // 出版社
		Date   string  `form:"date" json:"date" binding:"omitempty,date"`                                                 // 出版日期
		Kind   uint8   `form:"kind" json:"kind" binding:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19"` // 分类
	}

	// BookUrlRequest url参数
	BookUrlRequest struct {
		ID int `uri:"id" json:"id" binding:"required,gt=0"`
	}
)
