package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	prev := &ListNode{Val: 0}
	current := prev

	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10
		current.Next = &ListNode{Val: sum}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		current.Next = &ListNode{Val: carry}
	}

	return prev.Next
}

func main() {

}
