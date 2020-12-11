/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	cur := head
	next := cur.Next

	for cur.Val != nil && next != nil {
		if cur.Val == cur.Next.Val {
			tmp := next.Next
			for tmp != nil && tmp.Val == cur.Val {
				tmp = tmp.Next
			}
			cur = tmp
		} else {
			cur = cur.Next
		}

		if cur = nil {
			return nil
		}

		next = cur.Next
	}
}
// @lc code=end

