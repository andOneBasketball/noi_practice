package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	arr1, arr2 := make([]int, 0, 32), make([]int, 0, 32)
	l := l1
	for l != nil {
		arr1 = append(arr1, l.Val)
		l = l.Next
	}

	l = l2
	for l != nil {
		arr2 = append(arr2, l.Val)
		l = l.Next
	}

	arr := make([]int, 0, 32)
	arr1Len, arr2Len := len(arr1), len(arr2)
	maxLen := arr1Len
	if arr2Len > maxLen {
		maxLen = arr2Len
	}
	i, j := 0, 0
	num := 0
	for i < maxLen {
		if i < arr1Len && j < arr2Len {
			mod := (arr1[i] + arr2[j] + num) % 10
			num = (arr1[i] + arr2[j] + num) / 10
			arr = append(arr, mod)
		} else if i < arr1Len {
			mod := (arr1[i] + num) % 10
			num = (arr1[i] + num) / 10
			arr = append(arr, mod)
		} else if i < arr2Len {
			mod := (arr2[i] + num) % 10
			num = (arr2[i] + num) / 10
			arr = append(arr, mod)
		}
		i++
		j++
	}
	if num > 0 {
		arr = append(arr, num)
	}
	//fmt.Println(arr1, arr2, arr)

	lNew := new(ListNode)
	lNew.Val = arr[0]
	ll := lNew
	for i := 1; i < len(arr); i++ {
		lTmp := new(ListNode)
		lTmp.Val = arr[i]
		ll.Next = lTmp
		ll = lTmp
	}

	return lNew
}

func main() {
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l := addTwoNumbers(l1, l2)
	for l.Next != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Printf("%d\n", l.Val)
}
