package main

import (
	"fmt"
)

func numIslandsDfs(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, x, y int) {
	m, n := len(grid), len(grid[0])
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
}

func numIslands(grid [][]byte) int {
	count := 0
	steps := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	visited := make(map[[2]int]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[[2]int{i, j}] {
				continue
			}

			stack := make([][2]int, 0, len(grid[i]))
			if grid[i][j] == '1' {
				stack = append(stack, [2]int{i, j})
				for len(stack) > 0 {
					point := stack[len(stack)-1]
					stack = stack[0 : len(stack)-1]
					visited[point] = true
					for _, step := range steps {
						x, y := point[0]+step[0], point[1]+step[1]
						if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) &&
							!visited[[2]int{x, y}] && grid[x][y] == '1' {
							stack = append(stack, [2]int{x, y})
						}
					}
				}
				count++
			}
		}
	}
	return count
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
	grid = [][]byte{
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
