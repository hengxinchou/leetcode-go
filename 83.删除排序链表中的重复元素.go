/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
 
func deleteDuplicates(head *ListNode) *ListNode {
	// nums := map[int]int{}
	// cur := head

	// var prev *ListNode
	// for cur != nil {
	// 	if _, ok := nums[cur.Val]; ok {
	// 		cur = cur.Next
	// 		prev.Next = cur
	// 	} else {
	// 		nums[cur.Val] = 1
	// 		prev = cur
	// 		cur = cur.Next
	// 	}
	// }
	// return head
    cur := head
	for cur != nil && cur.Next != nil {
		// fmt.Printf("cur's value is %d\n", cur.Val)
		if cur.Val == cur.Next.Val {
			// fmt.Printf("if logic, cur's value is %d\n", cur.Val)
			cur.Next = cur.Next.Next
			// cur = cur.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
// @lc code=end
func print(head *ListNode){
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func main(){
	l11 := ListNode{Val: 1}
	l12 := ListNode{Val: 1}
	l13 := ListNode{Val: 2}
	l14 := ListNode{Val: 3}
	l15 := ListNode{Val: 3}
	l16 := ListNode{Val: 3}
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14
	l14.Next = &l15
	l15.Next = &l16

	print(&l11)
	l21 := deleteDuplicates(&l11)
	print(l21)
}
