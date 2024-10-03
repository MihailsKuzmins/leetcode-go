package leetcode_1

import "testing"

func Test1(t *testing.T) {
	nums, target := []int{2, 7, 11, 15}, 9
	expected := []int{0, 1}

	actual := twoSum(nums, target)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func Test2(t *testing.T) {
	nums, target := []int{3, 2, 4}, 6
	expected := []int{1, 2}

	actual := twoSum(nums, target)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func Test3(t *testing.T) {
	nums, target := []int{3, 3}, 6
	expected := []int{0, 1}

	actual := twoSum(nums, target)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual[i])
		}
	}
}
