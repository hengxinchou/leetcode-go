/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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
	v1 := &ListNode{Val: 1}
	v2 := &ListNode{Val: 2}
	v3 := &ListNode{Val: 3}
	v4 := &ListNode{Val: 4}
	v5 := &ListNode{Val: 5}
	v6 := &ListNode{Val: 6}
	v7 := &ListNode{Val: 7}
	v8 := &ListNode{Val: 8}
	v9 := &ListNode{Val: 9}

	v1.Next = v2
	v2.Next = v3
	v3.Next = v4
	v4.Next = v5
	v5.Next = v6
	v6.Next = v7
	v7.Next = v8
	v8.Next = v9
	v9.Next = v5

	// print(v1)
	tmp := detectCycle(v1)
	fmt.Println(tmp.Val)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1 := head
	p2 := head

	for p2.Next != nil && p2.Next.Next != nil {

		p1 = p1.Next
		p2 = p2.Next.Next
		if p2 == nil {
			return nil
		}

		if p1 == p2 {
			// fmt.Println("here")
			p := head
			for p != p1 {
				// fmt.Println("there")
				p = p.Next
				p1 = p1.Next
			}
			return p
		}
	}
	return nil
}

// @lc code=end
