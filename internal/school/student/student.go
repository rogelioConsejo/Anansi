package student

func NewStudent(name string) Student {
	return student{name}
}

type Student interface {
	Name() string
}

type student struct {
	name string
}

func (s student) Name() string {
	return s.name
}
