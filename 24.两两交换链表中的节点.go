/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Definition for singly-linked list.
package main
import "fmt"
type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	l1 := ListNode{Val: 1}
	// l2 := ListNode{Val: 2}
	// l3 := ListNode{Val: 3}
	// l4 := ListNode{Val: 4}
	// l5 := ListNode{Val: 5}

	// l1.Next = &l2
	// l2.Next = &l3
	// l3.Next = &l4
	// l4.Next = &l5

	print(&l1)
	tmp := swapPairs(&l1)

	print(tmp)
}


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &ListNode{}
	
	cur := head

	// var odd *ListNode
	var even *ListNode

	prev := newHead
	for cur != nil && cur.Next != nil {
		
		even = cur.Next
		nextPairHead := even.Next
		
		prev.Next = even
		even.Next = cur
		cur.Next = nextPairHead
		
		prev = cur
		cur = nextPairHead
		// if cur != nil {
		// 	fmt.Printf("prev is %d, cur is %d\n", prev.Val, cur.Val)
		// }
		
	}
	return newHead.Next
}


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}


func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
// @lc code=end

