package main

import "fmt"

type LRUCache2 struct {
	capacity int
	hash     map[int]*Node
	Head     *Node
	Tail     *Node
}

type Node struct {
	key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor2(capacity int) LRUCache2 {
	h := &Node{-1, -1, nil, nil}
	t := &Node{-1, -1, nil, nil}
	h.Next = t
	t.Next = h

	hash := make(map[int]*Node, capacity)
	cache := LRUCache2{hash: hash, capacity: capacity, Head: h, Tail: t}
	return cache
}

func (this *LRUCache2) insert(node *Node) {
	t := this.Tail
	node.Prev = t.Prev
	t.Prev.Next = node
	node.Next = t
	t.Prev = node
}

func (this *LRUCache2) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache2) Get(key int) int {

	if value, ok := this.hash[key]; ok {
		this.remove(value)
		this.insert(value)
		return value.Val
	} else {

		return -1
	}

}

func (this *LRUCache2) Put(key, value int) {

	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Val = value
	} else {
		if len(this.hash) >= this.capacity {
			h := this.Head.Next
			this.remove(h)
			delete(this.hash, h.key)
		}
		node := &Node{key: key, Val: value, Prev: nil, Next: nil}
		this.hash[key] = node
		this.insert(node)
	}
}

func main() {
	lru := Constructor2(4)

	lru.Put(1, 1)
	fmt.Println(lru.Get(1))

	lru.Put(2, 2)
	fmt.Println(lru.Get(2))

	lru.Put(3, 3)
	fmt.Println(lru.Get(3))

	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
	lru.Put(5, 5)
	lru.Put(6, 6)
	fmt.Println(lru.hash)
}
