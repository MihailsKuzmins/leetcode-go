package leetcode_3

import "testing"

func Test1(t *testing.T) {
	s := "abcabcbb"
	expected := 3

	actual := lengthOfLongestSubstring(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test2(t *testing.T) {
	s := "bbbbb"
	expected := 1

	actual := lengthOfLongestSubstring(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test3(t *testing.T) {
	s := "pwwkew"
	expected := 3

	actual := lengthOfLongestSubstring(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}

func Test4(t *testing.T) {
	s := "abba"
	expected := 2

	actual := lengthOfLongestSubstring(s)

	if actual != expected {
		t.Errorf("Expected to be %d, but got %d", expected, actual)
	}
}
