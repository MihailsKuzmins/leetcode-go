package leetcode_1963

import "testing"

func Test1(t *testing.T) {
	s := "][]["
	expected := 1

	actual := minSwaps(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test2(t *testing.T) {
	s := "]]][[["
	expected := 2

	actual := minSwaps(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test3(t *testing.T) {
	s := "[]"
	expected := 0

	actual := minSwaps(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test4(t *testing.T) {
	s := "[[][][]]]["
	expected := 1

	actual := minSwaps(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}
