package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	L        *list.List
	Data     map[int]*list.Element
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		L:        list.New(),
		Data:     make(map[int]*list.Element, capacity),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	result, ok := this.Data[key]
	if !ok {
		return -1
	}
	this.L.MoveToFront(result)
	return result.Value.([2]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		node.Value = [2]int{key, value}
		this.L.MoveToFront(node)
	} else {
		if this.L.Len() == this.Capacity {
			oldNode := this.L.Back()
			delete(this.Data, oldNode.Value.([2]int)[0])
			this.L.Remove(oldNode)
		}
		node = this.L.PushFront([2]int{key, value})
		this.Data[key] = node
	}
}

func (this *LRUCache) PrintList() {
	for l := this.L.Front(); l != nil; l = l.Next() {
		fmt.Printf("%v ", l.Value)
	}
	fmt.Println()
}

func main() {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4), cache.Data)
	fmt.Println(cache.Get(3), cache.Data)
	fmt.Println(cache.Get(2), cache.Data)
	cache.PrintList()
	fmt.Println(cache.Get(1), cache.Data)
	cache.Put(5, 5)
	fmt.Println(cache.Data)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(5))
}
