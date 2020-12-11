/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
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

func middleNode(head *ListNode) *ListNode {
	
	if head == nil {
		return nil
	}

	cur1 := head
	cur2 := head

	for cur2 != nil && cur2.Next != nil {
		cur1 = cur1.Next
		cur2 = cur1.Next.Next
	}
	return cur1

// 	i := 0
// 	cur := head
// 	for cur != nil {
// 		i++
// 		cur = cur.Next
// 	}

// 	j := i / 2

// 	// fmt.Printf("j is %d\n", j)
// 	cur = head
// 	for k := 0 ; k < j ; k++{
// 		cur = cur.Next
// 	}

// 	return cur
// }
// @lc code=end
}

func main(){
	l11 := ListNode{Val: 1}
	l12 := ListNode{Val: 2}
	l13 := ListNode{Val: 3}
	l14 := ListNode{Val: 4}
	l15 := ListNode{Val: 5}
	l16 := ListNode{Val: 6}
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14
	l14.Next = &l15
	l15.Next = &l16
	l111 := middleNode(&l11)
	fmt.Println(l111.Val)
}
