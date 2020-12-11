/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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


func isPalindrome(head *ListNode) bool {

	tmp := []int{}

	cur := head
	for cur != nil {
		tmp = append(tmp, cur.Val)
		cur = cur.Next
	}
	if len(tmp) == 1 {
		return true
	}
	mid := len(tmp) / 2
	for i := 0 ; i < mid ; i++ {
		if tmp[i] != tmp[len(tmp) - i -1]{
			return false
		} 
	}
	return true
}

func main() {
	
	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 2}
	l3 := ListNode{Val: 2}
	l4 := ListNode{Val: 1}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	fmt.Printf("%t\n", isPalindrome(&l1))
}

// @lc code=end

