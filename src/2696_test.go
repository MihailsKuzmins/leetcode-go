package leetcode_2696

import "testing"

func Test1(t *testing.T) {
	s := "ABFCACDB"
	expected := 2

	actual := minLength(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test2(t *testing.T) {
	s := "ACBBD"
	expected := 5

	actual := minLength(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test3(t *testing.T) {
	s := "ABBA"
	expected := 2

	actual := minLength(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test4(t *testing.T) {
	s := "CCCCDDDD"
	expected := 0

	actual := minLength(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test5(t *testing.T) {
	s := "CCDAABBDCD"
	expected := 0

	actual := minLength(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}
