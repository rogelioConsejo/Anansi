package course

func NewCourse(s string) Course {
	return &course{name: s}
}

type Course interface {
	Name() string
	AppendLesson(l Lesson)
	Lessons() []Lesson
}

type course struct {
	name    string
	lessons []Lesson
}

func (c *course) Lessons() []Lesson {
	return c.lessons
}

func (c *course) AppendLesson(l Lesson) {
	c.lessons = append(c.lessons, l)
}

func (c *course) Name() string {
	return c.name
}
