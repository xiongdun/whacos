package main

import "fmt"

func main() {
	s := "abcdef ghijk lmn opqrst"

	fmt.Println(reverseString(s))

}

func reverseString(s string) string {

	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//func mergeList(node1 *ListNode, node2 *ListNode) *ListNode {
//
//	if node1 == nil ||  node2 == nil{
//		return nil
//	}
//
//	var mergedHead *ListNode
//
//	if node1.Val < node2.Val {
//
//		mergedHead = node1
//		mergedHead = mergeList(node1.Next, node2)
//	} else {
//		mergedHead = node2
//		mergedHead = mergeList(node1, node2.Next)
//	}
//
//	return mergedHead
//}
