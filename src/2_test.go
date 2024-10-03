package leetcode_2

import "testing"

func Test1(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expected := []int{7, 0, 8}

	actual := addTwoNumbers(l1, l2)

	if actual != nil {
		for i, val := range expected {
			if actual.Val != val {
				t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual.Val)
			}

			actual = actual.Next
		}
	}
}

func Test2(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}
	expected := []int{0}

	actual := addTwoNumbers(l1, l2)

	if actual != nil {
		for i, val := range expected {
			if actual.Val != val {
				t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual.Val)
			}

			actual = actual.Next
		}
	}
}

func Test3(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	actual := addTwoNumbers(l1, l2)

	if actual != nil {
		for i, val := range expected {
			if actual.Val != val {
				t.Errorf("Expected an element at %d to be %d, but got %d", i, expected[i], actual.Val)
			}

			actual = actual.Next
		}
	}
}
