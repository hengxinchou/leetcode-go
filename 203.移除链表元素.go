/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
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
	// l2 := &ListNode{Val: 2}	
	// l3 := &ListNode{Val: 6}	
	// l4 := &ListNode{Val: 3}	
	// l5 := &ListNode{Val: 4}	
	// l6 := &ListNode{Val: 5}	
	// l7 := &ListNode{Val: 6}	

	// l1.Next = l2
	// l2.Next = l3
	// l3.Next = l4
	// l4.Next = l5
	// l5.Next = l6
	// l6.Next = l7
	print(l1)
	l1 = removeElements(l1, 1)
	print(l1)
}

func print(head *ListNode){
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}			
	cur := head
	newHead := &ListNode{}
	prev := newHead
	for cur != nil {
		if cur.Val == val {
			fmt.Println("hello")
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev.Next = cur
		    prev = cur
		    cur = cur.Next
		}
	}
	return newHead.Next
}
// @lc code=end

