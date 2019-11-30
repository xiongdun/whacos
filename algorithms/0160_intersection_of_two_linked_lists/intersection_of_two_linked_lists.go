package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for pA != pB {

		if pA == nil {
			pA = headB
		} else {
			pA = headA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = headB.Next
		}
	}

	return pA
}
