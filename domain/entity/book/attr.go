package book

type (
	Attribute  func(*Book)
	Attributes []Attribute
)

func (a Attributes) Apply(book *Book) {
	for _, attribute := range a {
		attribute(book)
	}
}

func WithBookID(id int) Attribute {
	return func(book *Book) {
		if id > 0 {
			book.ID = id
		}
	}
}

func WithBookName(name string) Attribute {
	return func(book *Book) {
		if name != "" {
			book.Name = name
		}
	}
}

func WithBookBlurb(blurb string) Attribute {
	return func(book *Book) {
		if blurb != "" {
			book.Blurb = blurb
		}
	}
}

func WithBookPrice1(price float64) Attribute {
	return func(book *Book) {
		if price >= 0 {
			book.Price1 = price
		}
	}
}

func WithBookPrice2(price float64) Attribute {
	return func(book *Book) {
		if price >= 0 {
			book.Price2 = price
		}
	}
}

func WithBookAuthor(author string) Attribute {
	return func(book *Book) {
		if author != "" {
			book.Author = author
		}
	}
}

func WithBookPress(press string) Attribute {
	return func(book *Book) {
		if press != "" {
			book.Press = press
		}
	}
}

func WithBookDate(date string) Attribute {
	return func(book *Book) {
		if date != "" {
			book.Date = date
		}
	}
}

func WithBookKind(kind uint8) Attribute {
	return func(book *Book) {
		if kind > 0 {
			book.Kind = kind
		}
	}
}

func WithBookKindStr(kindStr string) Attribute {
	return func(book *Book) {
		if kindStr != "" {
			book.KindStr = kindStr
		}
	}
}
