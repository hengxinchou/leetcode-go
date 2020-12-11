/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	curB := headB 

	flagA, flagB := true, true
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		if curA == nil && flagA {
			curA = headB
			flagA = false
		}
		curB = curB.Next
		if curB == nil && flagB {
			curB = headA
			flagB = false
		}
	}
	return nil
}
// @lc code=end

func print(l *ListNode){

	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}
func main(){
	l11 := ListNode{Val: 4}
	l12 := ListNode{Val: 1}
	l3 := ListNode{Val: 8}
	l4 := ListNode{Val: 4}
	l5 := ListNode{Val: 5}
	l11.Next = &l12
	l12.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	l21 := ListNode{Val: 5}
	// l22 := ListNode{Val: 0}
	// l23 := ListNode{Val: 1}

	// l21.Next = &l22
	// l22.Next = &l23
	// l23.Next = &l3

	print(&l11)
	print(&l21)
	tmp := getIntersectionNode(&l11, &l21)
	if tmp != nil {
		fmt.Println(tmp.Val)
	}
	
}

