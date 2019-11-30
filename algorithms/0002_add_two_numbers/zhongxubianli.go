package main

import (
	"fmt"
	"github.com/Unknwon/com"
)

type BinaryTreeNode struct {
	Value int
	Right *BinaryTreeNode
	Left  *BinaryTreeNode
}

func zhongxubianli(root *BinaryTreeNode, k int) *BinaryTreeNode {
	var target *BinaryTreeNode = nil

	if root.Left != nil {
		target = kthNode(root, k)
	}

	if target == nil {
		if k == 1 {
			target = root
		}
		k--
	}

	if target == nil && root.Right != nil {
		target = kthNode(root, k)
	}

	return target
}

func kthNode(root *BinaryTreeNode, k int) *BinaryTreeNode {

	if root == nil || k == 0 {
		return nil
	}

	return kthNode(root, k)
}

func main() {
	reverse := com.Reverse("xiongdun")
	fmt.Println(reverse)
}
