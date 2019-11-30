package main

import "fmt"

func main() {
	questions := []Question{
		{
			Parameter{[]int{}, []int{}},
			Answer{[]int{}},
		},

		{
			Parameter{[]int{1}, []int{1}},
			Answer{[]int{2}},
		},

		{
			Parameter{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			Answer{[]int{2, 4, 6, 8}},
		},

		{
			Parameter{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			Answer{[]int{2, 4, 6, 8, 0, 1}},
		},

		{
			Parameter{[]int{1}, []int{9, 9, 9, 9, 9}},
			Answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			Parameter{[]int{9, 9, 9, 9, 9}, []int{1}},
			Answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			Parameter{[]int{2, 4, 3}, []int{5, 6, 4}},
			Answer{[]int{7, 0, 8}},
		},

		{
			Parameter{[]int{1, 8, 3}, []int{7, 1}},
			Answer{[]int{8, 9, 3}},
		},
	}

	for _, question := range questions {
		_, param := question.Answer, question.Parameter
		fmt.Println(param)
	}
}

type Parameter struct {
	one     []int
	another []int
}

type Answer struct {
	one []int
}

type Question struct {
	Parameter
	Answer
}

type ListNode2 struct {
	Value int
	Next  *ListNode2
}

func addTwoNumbers2(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode2{Value: 0, Next: nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			///
			///
			///
			x = l1.Value
		}

		if l2 == nil {
			y = 0
		} else {
			//
			//
			//
			//
			fmt.Println("")
			y = l2.Value
		}

		current.Next = &ListNode2{Value: (x + y + carry) % 10, Next: nil}
		current = current.Next
		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head.Next
}
