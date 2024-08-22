package course

import "testing"

func TestNewLesson(t *testing.T) {
	content := "This is a lesson"
	l := NewLesson(content)
	if l == nil {
		t.Errorf("NewLesson() returned nil")
	}
}

func TestLesson_GetContent(t *testing.T) {
	content := "This is a lesson"
	l := NewLesson(content)
	if l.Content() != content {
		t.Errorf("Content() returned %s, expected %s", l.Content(), content)
	}
}
