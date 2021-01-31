/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
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
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	// 新建链表哨兵
	OrderedHead := &ListNode{Next: head}

	cur := head.Next
	head.Next = nil // 封闭上, 以防插入到已排序链表时，带上了后面的未排序链表
	// 遍历未排序链表
	for cur != nil {
		tmp := cur.Next // 先记住下一个要遍历的链点

		// 遍历已排序列表，寻找插入位置, 即第一个比待插入链点值大的位置
		pre := OrderedHead
		orderedNode := OrderedHead.Next
		for orderedNode != nil && orderedNode.Val <= cur.Val {
			pre = orderedNode
			orderedNode = orderedNode.Next
		}

		// 待插入位置是最后节点
		if orderedNode == nil {
			// 到了OrderedList的末尾节点
			// fmt.Printf("else loginc, pre val is %d\n", pre.Val)
			pre.Next = cur
			cur.Next = nil // 要封闭上
		} else {
			//待插入位置不是最后节点，而是中间节点
			pre.Next = cur
			cur.Next = orderedNode
		}

		// 遍历下一个待插入节点
		cur = tmp

	}
	return OrderedHead.Next
}

// @lc code=end

func print(head *ListNode) {
	fmt.Println("print the list")
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l11 := ListNode{Val: -1}
	l12 := ListNode{Val: 5}
	l13 := ListNode{Val: 3}
	l14 := ListNode{Val: 4}
	l15 := ListNode{Val: 0}

	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14
	l14.Next = &l15

	print(&l11)
	a := insertionSortList(&l11)
	print(a)

	l21 := ListNode{Val: 3}
	l22 := ListNode{Val: 4}
	l23 := ListNode{Val: 1}

	l21.Next = &l22
	l22.Next = &l23

	print(&l21)
	b := insertionSortList(&l21)
	print(b)

}
