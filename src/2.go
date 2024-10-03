package leetcode_2

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	remainder, ir := 0, result

	for {
		sum := remainder
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		ir.Val = sum % 10
		remainder = sum / 10

		if l1 == nil && l2 == nil {
			if remainder != 0 {
				ir.Next = &ListNode{Val: remainder}
			}

			break
		}

		ir.Next = &ListNode{}
		ir = ir.Next
	}

	return result
}
