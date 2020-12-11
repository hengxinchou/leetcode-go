/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
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
	l1 := &ListNode{Val: 4}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 3}
	l5 := &ListNode{Val: -1}
	l6 := &ListNode{Val: 5}
	l7 := &ListNode{Val: 2}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7

	print(l1)
	l11 := sortList(l1)	 
	print(l11)
}

func print(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	cur := head
	next := head.Next

	for next != nil && next.Next != nil {
		cur = cur.Next
		next = next.Next.Next
	}

	// fmt.Printf("cur is %d\n", cur.Val)

	l1 := head
	l2 := cur.Next
	cur.Next = nil

	sort1 := sortList(l1)
	sort2 := sortList(l2)

	res := merge(sort1, sort2) 
	return res
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil  {
		return l2
	}
	if l2 == nil {
		return l1
	}


	// var prev1, prev2, *ListNode
	cur1 := l1
	cur2 := l2

	newHead := &ListNode{}
	if cur1.Val <= cur2.Val {
		newHead.Next = cur1
		cur1 = cur1.Next
	} else {
		newHead.Next = cur2
		cur2 = cur2.Next
	}
	prev := newHead.Next
	// cur.next = nil

	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			prev.Next = cur1
			cur1 = cur1.Next
		} else {
			prev.Next = cur2
			cur2 = cur2.Next
		}
		prev = prev.Next
		// cur = cur.Next
		// cur.Next = nil
	}
	if cur2 != nil {
		prev.Next = cur2
	} 
	if cur1 != nil {
		prev.Next = cur1
	}
	return newHead.Next
}

func sortList1(head *ListNode) *ListNode {
	if head == nil {
        return head
    }
	sorted := head
	cur := head.Next
	sorted.Next = nil

	for cur != nil {
		var prevSorted, tmp *ListNode
		curSorted := sorted
		for curSorted != nil && cur.Val >= curSorted.Val {
			prevSorted = curSorted
			curSorted = curSorted.Next	
		}
		tmp = cur.Next
		if prevSorted == nil {
			cur.Next = curSorted
			sorted = cur
		} else {
			prevSorted.Next = cur
			cur.Next = curSorted
		}
		cur = tmp	
	}
	return sorted
}
// @lc code=end

