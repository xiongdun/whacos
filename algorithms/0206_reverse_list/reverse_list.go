package main

import "fmt"

func main() {
	fmt.Println(reverse(&ListNode{}))
}

type ListNode struct {
	next  *ListNode
	value int
}

// 反转链表，只需将指向下一个节点的指针指向前一个
func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var current = head
	var temp *ListNode = nil

	for current != nil {
		temp = current.next
		current.next = prev
		prev = current
		current = temp
	}
	return prev
}
