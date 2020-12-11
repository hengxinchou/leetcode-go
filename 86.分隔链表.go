/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
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
	Val  int
	Next *ListNode
}

func main() {
	v1 := ListNode{1, nil}
	v2 := ListNode{4, nil}
	v3 := ListNode{3, nil}
	v4 := ListNode{2, nil}
	v5 := ListNode{5, nil}
	v6 := ListNode{2, nil}
	v1.Next = &v2
	v2.Next = &v3
	v3.Next = &v4
	v4.Next = &v5
	v6.Next = &v6
	print(&v1)

	l1 := partition(&v1, 3)
	fmt.Println("partition list")
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

func partition(head *ListNode, x int) *ListNode {
	smallHead := &ListNode{}
	smallPrev := smallHead

	biggerHead := &ListNode{}
	biggerPrev := biggerHead

	cur := head
	for cur != nil {
		if cur.Val < x {
			smallPrev.Next = cur
			smallPrev = cur
			cur = cur.Next
		} else {
			biggerPrev.Next = cur
			biggerPrev = cur
			cur = cur.Next
		}
	}

	if smallPrev.Next != nil {
		smallPrev.Next = biggerHead.Next
		return smallHead.Next
	} else {
		return biggerHead.Next
	}
}

// @lc code=end
