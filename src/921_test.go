package leetcode_921

import "testing"

func Test1(t *testing.T) {
	s := "())"
	expected := 1

	actual := minAddToMakeValid(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test2(t *testing.T) {
	s := "((("
	expected := 3

	actual := minAddToMakeValid(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test3(t *testing.T) {
	s := ")))((("
	expected := 6

	actual := minAddToMakeValid(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}
