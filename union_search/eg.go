package main

import "fmt"

// 并查集结构体
type UnionFind struct {
	parent []int
	rank   []int
}

// 初始化并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i // 每个节点初始是自己的父节点
		rank[i] = 1   // 初始秩为1
	}
	return &UnionFind{parent, rank}
}

// 查找根节点，带路径压缩
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// 合并两个集合，按秩合并
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return // 已在一个集合
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// 判断两个元素是否在同一集合
func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main() {
	uf := NewUnionFind(10)

	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(4, 5)

	fmt.Println(uf.IsConnected(1, 3)) // true
	fmt.Println(uf.IsConnected(1, 5)) // false

	uf.Union(3, 5)
	fmt.Println(uf.IsConnected(1, 5)) // true
}
