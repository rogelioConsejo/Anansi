package syllabus

type Class interface {
	GetContent() Content
	GetTitle() string
}

type Content interface {
}
