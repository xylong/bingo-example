package answer

type (
	Attribute  func(*Answer)
	Attributes []Attribute
)

// Apply 应用属性
func (a Attributes) Apply(answer *Answer) {
	for _, attribute := range a {
		attribute(answer)
	}
}

func WithID(id int) Attribute {
	return func(answer *Answer) {
		if id > 0 {
			answer.ID = id
		}
	}
}

func WithQuestionID(id int) Attribute {
	return func(answer *Answer) {
		if id > 0 {
			answer.QuestionID = id
		}
	}
}

func WithContent(content string) Attribute {
	return func(answer *Answer) {
		if content != "" {
			answer.Content = content
		}
	}
}

func WithIsCorrect(isCorrect bool) Attribute {
	return func(answer *Answer) {
		answer.IsCorrect = &isCorrect
	}
}

func WithOther(other string) Attribute {
	return func(answer *Answer) {
		if other != "" {
			answer.Other = other
		}
	}
}

func WithImg(img string) Attribute {
	return func(answer *Answer) {
		if img != "" {
			answer.Img = img
		}
	}
}
