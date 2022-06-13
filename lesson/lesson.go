package lesson

type Lesson interface {
	Show() Content
}

type Content string

type lesson struct {
}

func (l lesson) Show() Content {
	return ""
}

func New() Lesson {
	return lesson{}
}
