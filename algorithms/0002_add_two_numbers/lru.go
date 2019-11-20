package main

import "fmt"

type LRUCache struct {
	m        map[int]int
	s        []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]int, capacity)
	s := make([]int, capacity)
	return LRUCache{m, s, capacity}
}

func (this *LRUCache) Get(key int) int {
	return this.m[key]
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.m) == this.capacity {

		dkey := this.s[0]

		delete(this.m, this.m[dkey])
	}

	this.m[key] = value

	for i := 0; i < this.capacity; i++ {
		if i == this.capacity-1 {
			this.s[i] = key
			break
		}
		this.s[i] = this.s[i+1]
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lru := Constructor(4)

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
	fmt.Println(lru.m)
}
