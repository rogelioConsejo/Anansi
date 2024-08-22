package student

import "testing"

func TestNewStudent(t *testing.T) {
	const studentName = "Test Student"
	s := NewStudent(studentName)
	if s == nil {
		t.Fatalf("NewStudent() returned nil")
	}

	if s.Name() != studentName {
		t.Errorf("student name is %s, expected %s", s.Name(), studentName)
	}
}
