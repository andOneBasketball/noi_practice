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
func reverseList(head *ListNode) *ListNode {
	arr := make([]int, 0, 32)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	fmt.Println(arr)
	if len(arr) == 0 {
		return nil
	}
	l := new(ListNode)
	tmp := l
	for i := len(arr) - 1; i >= 0; i-- {
		tmp.Val = arr[i]
		if i > 0 {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		}
	}
	return l
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	ll := reverseList(l)
	for ll != nil {
		fmt.Printf("%d ", ll.Val)
		ll = ll.Next
	}
}
