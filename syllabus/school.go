package syllabus

type School interface {
	GetCourses() map[string]Course
}

type Class interface {
	GetContent() Content
	GetTitle() string
}

type Content interface {
}
