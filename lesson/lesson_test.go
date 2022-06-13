package lesson

import "testing"

func TestShowLesson(t *testing.T) {
	var testLesson Lesson = New()
	var _ Content = testLesson.Show()
}
