package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (this *Node) Constructor(value int, next *Node, random *Node) {
	this.Val = value
	this.Next = next
	this.Random = random
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	pointer := head

	for pointer != nil {
		newNode := &Node{Val: pointer.Val}

		newNode.Next = pointer.Next

		pointer.Next = newNode
		pointer = newNode.Next
	}

	pointer = head

	for pointer != nil {
		if pointer.Random != nil {
			pointer.Next.Random = pointer.Random.Next
		} else {
			pointer.Next.Random = nil
		}

		pointer = pointer.Next.Next
	}

	pointerOldList := head
	pointerNewList := head.Next
	headOld := head.Next

	for pointerOldList != nil {
		pointerOldList.Next = pointerOldList.Next.Next

		if pointerNewList.Next != nil {
			pointerNewList.Next = pointerNewList.Next.Next
		} else {
			pointerNewList.Next = nil
		}

		pointerOldList = pointerOldList.Next
		pointerNewList = pointerNewList.Next
	}

	return headOld
}

func main() {
	fmt.Println(copyRandomList(&Node{}))
}
