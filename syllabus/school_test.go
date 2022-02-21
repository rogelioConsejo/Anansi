package syllabus

import "testing"

func TestGetClasses(t *testing.T) {
	var TestSchool School = GetTestSchool("test")
	testCourses := TestSchool.GetCourses()
	testCourses[testCourseName].GetClasses()
}

type testSchool struct {
}

type testCourse struct {
}

func (t testSchool) GetCourses() map[string]Course {
	courses := make(map[string]Course)
	courses[testCourseName] = MakeCourse(testCourseName)
	return courses
}

func GetTestSchool(s string) School {
	return new(testSchool)
}

const testCourseName = "test_course"
