package entity

const (
	OrderByPriceAsc  = iota + 1 // 价格从低到高
	OrderByPriceDesc            // 价格从高到低
)

type Books []*Book

type Book struct {
	BookID     int     `gorm:"column:book_id;AUTO_INCREMENT;PRIMARY_KEY" json:"book_id"`
	BookName   string  `gorm:"column:book_name;type:varchar(50)" json:"book_name"`
	BookIntr   string  `gorm:"column:book_intr;type:text" json:"book_intr"`
	BookPrice1 float64 `gorm:"column:book_price1;type:decimal" json:"book_price1"`
	BookPrice2 float64 `gorm:"column:book_price2;type:decimal" json:"book_price2"`
	BookAuthor string  `gorm:"column:book_author;type:varchar(50)" json:"book_author"`
	BookPress  string  `gorm:"column:book_press;type:varchar(50)" json:"book_press"`
	BookDate   string  `gorm:"column:book_date;type:varchar(50)" json:"book_date"`
	BookKind   int     `gorm:"column:book_kind;type:int" json:"book_kind"`
}
