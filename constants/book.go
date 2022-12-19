package constants

// 字段
const (
	BookIndex  = "books"  // 📚es索引
	BookName   = "name"   // 书名
	BookPress  = "press"  // 出版社
	BookBlurb  = "blurb"  // 简介
	BookPrice1 = "price1" // 价格1
	BookPrice2 = "price2" // 价格2
	BookDate   = "date"   //出版日期
)

// 排序
const (
	BookPrice1Desc = iota + 1 // 价格1降序
	BookPrice1Asc             // 价格1升序
	BookPrice2Desc            // 价格2降序
	BookPrice2Asc             // 价格2升序
	BookDateDesc              // 出版日期降序
	BookDateAsc               // 出版日期升序
)

const (
	// BookRequest graphql请求字符串
	BookRequest = `
query{
	Book{
		id
		name
	}
}
`
)
