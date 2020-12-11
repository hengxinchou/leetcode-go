/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	// l4 := &ListNode{Val: 4}
	// l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	// l3.Next = l4
	// l4.Next = l5

	print(l1)
	// l11 := reverseList(l1)
	// print(l11)

	reorderList(l1)
	print(l1)
}

func print(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()	
}



func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	cur := head	
	next := head
	for next != nil && next.Next != nil {
		cur = cur.Next
		next = next.Next.Next
	}

	newHead := reverseList(cur.Next)
	cur.Next = nil

	cur1 := head
	cur2 := newHead

	var tmp1, tmp2 *ListNode

	for cur1 != nil && cur2 != nil {
		tmp1 = cur1.Next
		tmp2 = cur2.Next
		cur1.Next = cur2
		cur2.Next = tmp1
		cur1 = tmp1
		cur2 = tmp2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil	
	var tmp *ListNode
	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

// @lc code=end

