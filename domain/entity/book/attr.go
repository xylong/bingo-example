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

func WithID(id int) Attribute {
	return func(book *Book) {
		if id > 0 {
			book.ID = id
		}
	}
}

func WithName(name string) Attribute {
	return func(book *Book) {
		if name != "" {
			book.Name = name
		}
	}
}

func WithBlurb(blurb string) Attribute {
	return func(book *Book) {
		if blurb != "" {
			book.Blurb = blurb
		}
	}
}

func WithPrice1(price float64) Attribute {
	return func(book *Book) {
		if price >= 0 {
			book.Price1 = price
		}
	}
}

func WithPrice2(price float64) Attribute {
	return func(book *Book) {
		if price >= 0 {
			book.Price2 = price
		}
	}
}

func WithAuthor(author string) Attribute {
	return func(book *Book) {
		if author != "" {
			book.Author = author
		}
	}
}

func WithPress(press string) Attribute {
	return func(book *Book) {
		if press != "" {
			book.Press = press
		}
	}
}

func WithDate(date string) Attribute {
	return func(book *Book) {
		if date != "" {
			book.Date = date
		}
	}
}

func WithKind(kind uint8) Attribute {
	return func(book *Book) {
		if kind > 0 {
			book.Kind = kind
		}
	}
}

func WithKindStr(kindStr string) Attribute {
	return func(book *Book) {
		if kindStr != "" {
			book.KindStr = kindStr
		}
	}
}
