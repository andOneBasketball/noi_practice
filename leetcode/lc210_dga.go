package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	allTask := make(map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		allTask[i] = struct{}{}
	}
	// 节点入度
	taskCounts := make(map[int]int, numCourses)
	taskDirects := make(map[int][]int, numCourses)
	for _, pre := range prerequisites {
		taskCounts[pre[0]]++
		if _, ok := taskDirects[pre[1]]; !ok {
			taskDirects[pre[1]] = make([]int, 0, numCourses)
		}
		taskDirects[pre[1]] = append(taskDirects[pre[1]], pre[0])
	}
	result := make([]int, 0, numCourses)
	for len(allTask) > 0 {
		tasks := []int{}
		for t := range allTask {
			if taskCounts[t] == 0 {
				tasks = append(tasks, t)
			}
		}
		for _, tt := range tasks {
			if _, ok := taskDirects[tt]; ok {
				for _, d := range taskDirects[tt] {
					if taskCounts[d] > 0 {
						taskCounts[d]--
					}
				}
			}
			delete(allTask, tt)
		}
		if len(tasks) == 0 {
			return []int{}
		}
		result = append(result, tasks...)
	}
	return result
}

func main() {
	fmt.Println(findOrder(2, [][]int{[]int{1, 0}}))
}
