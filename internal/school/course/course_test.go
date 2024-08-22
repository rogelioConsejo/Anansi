package course

import "testing"

func TestNewCourse(t *testing.T) {
	course := NewCourse("Test Course")
	if course.Name() != "Test Course" {
		t.Errorf("Expected course name to be 'Test Course', got %s", course.Name())
	}
}

func TestCourse_AppendLesson(t *testing.T) {
	const courseName = "Test Course"
	const lessonContent = "Test Lesson"
	c := NewCourse(courseName)
	l := NewLesson(lessonContent)
	c.AppendLesson(l)
	if len(c.Lessons()) != 1 {
		t.Errorf("Expected 1 lesson, got %d", len(c.Lessons()))
	}
	if c.Lessons()[0].Content() != lessonContent {
		t.Errorf("Expected lesson content to be %s, got %s", lessonContent, c.Lessons()[0].Content())
	}

	const secondLessonContent = "Second Lesson"
	l2 := NewLesson(secondLessonContent)
	c.AppendLesson(l2)
	if len(c.Lessons()) != 2 {
		t.Errorf("Expected 2 lessons, got %d", len(c.Lessons()))
	}
}
