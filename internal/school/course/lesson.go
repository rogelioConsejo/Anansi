package course

func NewLesson(content string) Lesson {
	return lesson{content}
}

type Lesson interface {
	Content() string
}

type lesson struct {
	content string
}

func (l lesson) Content() string {
	return l.content
}
