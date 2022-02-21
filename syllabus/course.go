package syllabus

type Course interface {
	GetClasses() map[string]Class
}

type course struct {
	name    string
	classes map[string]Class
}

func (c course) GetClasses() map[string]Class {
	return c.classes
}

func MakeCourse(name string) Course {
	course := new(course)
	course.name = name
	course.classes = make(map[string]Class)

	return course
}
