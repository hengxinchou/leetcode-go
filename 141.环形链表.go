/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p1 := head
	p2 := head

	// fmt.Printf("p2.Next is %d\n", p2.Next.Val)

	for p1 != nil && p2 != nil && p2.Next != nil {

		fmt.Printf("p1 is %d\n", p1.Val)
		fmt.Printf("p2 is %d\n", p2.Val)
		fmt.Printf("p2.Next is %d\n", p2.Next.Val)
		
		p1 = p1.Next
		p2 = p2.Next.Next
		// fmt.Printf("p1 is %d\n", p2.Val)
		if p2 == nil {
			return false
		}

		if p1.Val == p2.Val {
			return true
		}
	}

	return false
}

func main() {
	 var result bool

	 v1 := &ListNode{0, nil}
	 v2 := &ListNode{1, nil}

	 v1.Next = nil
	 result = hasCycle(v1)
	 fmt.Printf("one node is %t\n", result)

	 v1.Next = v2
	 result = hasCycle(v1)
	 fmt.Printf("two node is %t\n", result)

	 fmt.Println("v2")
	 v21 := &ListNode{0, nil}
	 v21.Next = v21
	 result = hasCycle(v21)
	 fmt.Printf("one circle node is %t\n", result)

	 fmt.Println("v3")
	 v31 := &ListNode{0, nil}
	 v32 := &ListNode{0, nil}
	 v31.Next = v32
	 v32.Next = v31
	 result = hasCycle(v31)
	 fmt.Printf("two circle node is %t\n", result)

}

// @lc code=end

