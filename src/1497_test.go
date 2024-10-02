package leetcode_1497

import "testing"

func Test1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	const k = 5

	actual := canArrange(arr, k)

	if !actual {
		t.Errorf("Expected to be true, but got %t", actual)
	}
}

func Test2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	const k = 7

	actual := canArrange(arr, k)

	if !actual {
		t.Errorf("Expected to be true, but got %t", actual)
	}
}

func Test3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	const k = 10

	actual := canArrange(arr, k)

	if actual {
		t.Errorf("Expected to be false, but got %t", actual)
	}
}

func Test4(t *testing.T) {
	arr := []int{-10, 10}
	const k = 2

	actual := canArrange(arr, k)

	if !actual {
		t.Errorf("Expected to be false, but got %t", actual)
	}
}

func Test5(t *testing.T) {
	arr := []int{-4, -7, 5, 2, 9, 1, 10, 4, -8, -3}
	const k = 3

	actual := canArrange(arr, k)

	if !actual {
		t.Errorf("Expected to be false, but got %t", actual)
	}
}
