/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 */

 package main
 import "fmt"

 type ListNode struct {
      Val int
      Next *ListNode
  }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur1 := l1
	cur2 := l2

	cur := &ListNode{}
	// if cur1.Val <= cur2.Val {
	// 	cur = cur1
	// 	cur1 = cur1.Next
	// } else {
	// 	cur = cur2
	// 	cur2 = cur2.Next
	// }

	head := cur
	// printList(cur1)
	// printList(cur2)

	for cur1 != nil && cur2 != nil {
		// fmt.Printf("value is %d, %d\n", cur1.Val, cur2.Val)
		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			// cur = cur1
			cur1 = cur1.Next
			printList(head)
		} else {
			cur.Next = cur2
			// cur = cur2
			cur2 = cur2.Next
			// printList(head)
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}

	return head.Next
}

func printList(l1 *ListNode) {
	cur := l1
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
// @lc code=end


func main (){

	l11 := ListNode{Val: 1}
	l12 := ListNode{Val: 2}
	l13 := ListNode{Val: 4}
	l11.Next = &l12
	l12.Next = &l13

	l21 := ListNode{Val: 1}
	l22 := ListNode{Val: 3}
	l23 := ListNode{Val: 4}
	l21.Next = &l22
	l22.Next = &l23

	// printList(&l11)
	// printList(&l21)

	l3 := mergeTwoLists(&l11, &l21)

	printList(l3)
}