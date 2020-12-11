/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var pre *Node = nil
		for i := 0 ; i < size; i ++ {
			cur := queue[0]
			queue = queue[1:len(queue)]
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return root	
}
// @lc code=end

