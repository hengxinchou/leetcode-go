/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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



func main() {
	v1 := ListNode{ 1, nil}
	v2 := ListNode{ 2, nil}
	v3 := ListNode{ 3, nil}
	v4 := ListNode{ 4, nil}
	v5 := ListNode{ 5, nil}
	v1.Next = &v2
	v2.Next = &v3
	v3.Next = &v4
	v4.Next = &v5
	print(&v1)

	l1 := reverseBetween(&v1, 2, 4)
	fmt.Println("reverse list")
	print(l1)
}

func print(head *ListNode) {
	if head == nil {
		return
	}
	var current *ListNode = head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	cur := head

	count := 1
	var prev1 *ListNode
	for cur != nil && count >= m {	
		prev1 = cur
		cur = cur.Next
		count++
	}

	tmp := cur // just mark a place

	var next *ListNode
	var prev2 *ListNode = nil
	count--
	for cur != nil && count >= n {
		next = cur.Next
		cur.Next = prev2
		prev2 = cur
		cur = next
		count++
	}

	print(prev2)
	
	if m > 1 {
		prev1.Next = prev2
		tmp.Next = cur
		return head
	} else {
		return prev2
	}
}
// @lc code=end

