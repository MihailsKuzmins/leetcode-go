package leetcode_2491

import "testing"

func Test1(t *testing.T) {
	skill := []int{3, 2, 5, 1, 3, 4}
	var expected int64 = 22

	actual := dividePlayers(skill)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test2(t *testing.T) {
	skill := []int{3, 4}
	var expected int64 = 12

	actual := dividePlayers(skill)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test3(t *testing.T) {
	skill := []int{1, 1, 2, 3}
	var expected int64 = -1

	actual := dividePlayers(skill)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test4(t *testing.T) {
	skill := []int{2, 3, 4, 2, 5, 5}
	var expected int64 = 32

	actual := dividePlayers(skill)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}
