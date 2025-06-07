package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	Freq     map[int]*list.List
	Data     map[int]*list.Element
	MinFreq  int
	Capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Freq:     make(map[int]*list.List, capacity),
		Data:     make(map[int]*list.Element, capacity),
		MinFreq:  0,
		Capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1
	}
	freq := node.Value.([3]int)[2]
	this.Freq[freq].Remove(node)
	if freq == this.MinFreq && this.Freq[freq].Len() == 0 {
		this.MinFreq++
	}
	freq++
	if this.Freq[freq] == nil {
		this.Freq[freq] = list.New()
	}
	// newNode := this.Freq[freq].PushFront(node.Value)  // 访问频率修改了，为啥 first code 的时候就这样直接插入了呢？...
	newNode := this.Freq[freq].PushFront([3]int{node.Value.([3]int)[0], key, freq})
	this.Data[key] = newNode
	return node.Value.([3]int)[0]
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.Data[key]; ok {
		freq := node.Value.([3]int)[2]
		this.Freq[freq].Remove(node)
		if freq == this.MinFreq && this.Freq[freq].Len() == 0 {
			this.MinFreq++
		}
		freq++
		if this.Freq[freq] == nil {
			this.Freq[freq] = list.New()
		}
		newNode := this.Freq[freq].PushFront([3]int{value, key, freq})
		this.Data[key] = newNode
	} else {
		if len(this.Data) == this.Capacity {
			l := this.Freq[this.MinFreq]
			delete(this.Data, l.Back().Value.([3]int)[1])
			l.Remove(l.Back())
		}
		if this.Freq[0] == nil {
			this.Freq[0] = list.New()
		}
		l := this.Freq[0]
		node := l.PushFront([3]int{value, key, 0})
		this.MinFreq = 0
		this.Data[key] = node
	}
}

func (this *LFUCache) PrintList() {
	for k, v := range this.Freq {
		fmt.Printf(" %d: ", k)
		for node := v.Front(); node != nil; node = node.Next() {
			fmt.Printf("%v ", node.Value)
		}
	}
	fmt.Println()
}

func main() {
	data := [][]int{[]int{10, 13}, []int{3, 17}, []int{6, 11}, []int{10, 5}, []int{9, 10}, []int{13}, []int{2, 19}, []int{2}, []int{3}, []int{5, 25}, []int{8}, []int{9, 22}, []int{5, 5}, []int{1, 30}, []int{11}, []int{9, 12}, []int{7}, []int{5}, []int{8}, []int{9}, []int{4, 30}, []int{9, 3}, []int{9}, []int{10}, []int{10}, []int{6, 14}, []int{3, 1}, []int{3}, []int{10, 11}, []int{8}, []int{2, 14}, []int{1}, []int{5}, []int{4}, []int{11, 4}, []int{12, 24}, []int{5, 18}, []int{13}, []int{7, 23}, []int{8}, []int{12}, []int{3, 27}, []int{2, 12}, []int{5}, []int{2, 9}, []int{13, 4}, []int{8, 18}, []int{1, 7}, []int{6}, []int{9, 29}, []int{8, 21}, []int{5}, []int{6, 30}, []int{1, 12}, []int{10}, []int{4, 15}, []int{7, 22}, []int{11, 26}, []int{8, 17}, []int{9, 29}, []int{5}, []int{3, 4}, []int{11, 30}, []int{12}, []int{4, 29}, []int{3}, []int{9}, []int{6}, []int{3, 4}, []int{1}, []int{10}, []int{3, 29}, []int{10, 28}, []int{1, 20}, []int{11, 13}, []int{3}, []int{3, 12}, []int{3, 8}, []int{10, 9}, []int{3, 26}, []int{8}, []int{7}, []int{5}, []int{13, 17}, []int{2, 27}, []int{11, 15}, []int{12}, []int{9, 19}, []int{2, 15}, []int{3, 16}, []int{1}, []int{12, 17}, []int{9, 1}, []int{6, 19}, []int{4}, []int{5}, []int{5}, []int{8, 1}, []int{11, 7}, []int{5, 2}, []int{9, 28}, []int{1}, []int{2, 2}, []int{7, 4}, []int{4, 22}, []int{7, 24}, []int{9, 26}, []int{13, 28}, []int{11, 26}}
	cache := Constructor(10)
	for _, v := range data {
		if len(v) == 1 {
			fmt.Printf("key: %d, value: %d,", v[0], cache.Get(v[0]))
		} else if len(v) == 2 {
			cache.Put(v[0], v[1])
			fmt.Println()
			fmt.Println("put:", v[0], v[1])
			cache.PrintList()
		}
	}
}
