package leetcode_1331

import (
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{40, 10, 20, 30}
	expected := []int{4, 1, 2, 3}

	actual := arrayRankTransform(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func Test2(t *testing.T) {
	arr := []int{100, 100, 100}
	expected := []int{1, 1, 1}

	actual := arrayRankTransform(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func Test3(t *testing.T) {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	expected := []int{5, 3, 4, 2, 8, 6, 7, 1, 3}

	actual := arrayRankTransform(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func Test4(t *testing.T) {
	arr := []int{14, -19, 12, -25, 34, -27, -48, -37, 14, -47, 40, 23, 46, -39, 48, -41, 18, -27, -4}
	expected := []int{11, 8, 10, 7, 14, 6, 1, 5, 11, 2, 15, 13, 16, 4, 17, 3, 12, 6, 9}

	actual := arrayRankTransform(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}
