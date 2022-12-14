package dto

type (
	// BookSearchParam ðæ¥è¯¢åæ°
	BookSearchParam struct {
		*Pagination
		Name    string  `form:"name" json:"name" binding:"omitempty"`                                       // ä¹¦å
		Press   string  `form:"press" json:"press" binding:"omitempty"`                                     // åºçç¤¾
		Lowest  float64 `form:"lowest" json:"lowest" binding:"omitempty,gte=0,lte=10000"`                   // æä½ä»·
		Highest float64 `form:"highest" json:"highest" binding:"omitempty,gte=0,lte=10000,gtefield=Lowest"` // æé«ä»·
		Sorts   string  `form:"sorts" json:"sorts" binding:"omitempty"`                                     // æåº
	}

	// BookStoreParam ðåå»ºåæ°
	BookStoreParam struct {
		Name   string  `form:"name" json:"name" binding:"required"`
		Blurb  string  `form:"blurb" json:"blurb" binding:"omitempty"` // ç®ä»
		Price1 float64 `form:"price1" json:"price1" binding:"omitempty,gte=0"`
		Price2 float64 `form:"price2" json:"price2" binding:"omitempty,gte=0"`
		Author string  `form:"author" json:"author" binding:"omitempty"`                                                  // ä½è
		Press  string  `form:"press" json:"press" binding:"omitempty"`                                                    // åºçç¤¾
		Date   string  `form:"date" json:"date" binding:"omitempty,date"`                                                 // åºçæ¥æ
		Kind   uint8   `form:"kind" json:"kind" binding:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19"` // åç±»
	}

	// BookUrlRequest urlåæ°
	BookUrlRequest struct {
		ID int `uri:"id" json:"id" binding:"required,gt=0"`
	}
)
