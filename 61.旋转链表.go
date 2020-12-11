/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
	// v4 := ListNode{ 4, nil}
	// v5 := ListNode{ 5, nil}
	v1.Next = &v2
	v2.Next = &v3
	// v3.Next = &v4
	// v4.Next = &v5
	print(&v1)

	l1 := rotateRight(&v1, 1)
	fmt.Println("rotateRight list")
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


func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	length := 0 
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		prev = cur
		cur = cur.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	prev.Next = head
	steps := length - k

	// fmt.Printf("steps is %d\n", steps)
	start := 1
	cur = head
	for start < steps {
		cur = cur.Next
		start++
	}
	newHead := cur.Next
	cur.Next = nil

	return newHead
}
// @lc code=end

