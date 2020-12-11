/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var next *ListNode
	current := head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		
	}
	return prev
}

func main (){
	v1 := &ListNode{ 1, nil}
	v2 := &ListNode{ 2, nil}
	v3 := &ListNode{ 3, nil}
	v1.Next = v2
	v2.Next = v3
	print(v1)

	l1 := reverseList(v1)
	fmt.Println()
	fmt.Println("reverse list")
	print(l1)

	fmt.Println("nil list reverse")
	l2 := reverseList(nil)
	print(l2)
	
	fmt.Println("tow nodes' list reverse")
	v21 := &ListNode{ 1, nil}
	v22 := &ListNode{ 2, nil}
	v21.Next = v22
	l2Resutl := reverseList(v21)
	print(l2Resutl)

	fmt.Println("one node's list reverse")
	v31 := &ListNode{1,nil}
	l3Result := reverseList(v31)
	print(l3Result)
}


