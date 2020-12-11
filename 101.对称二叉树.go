/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main
import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && check(a.Left, b.Right) && check(a.Right, b.Left)
}
// @lc code=end

func print(root *TreeNode){
	fmt.Printf("%d ", root.Val)
	if root.Left != nil {
		print(root.Left)
	}
	if root.Right != nil {
		print(root.Right)
	}
}

func main() {
	t11 := TreeNode{Val: 1}
	t21 := TreeNode{Val: 2}
	t22 := TreeNode{Val: 2}
	t31 := TreeNode{Val: 2}
	// t32 := TreeNode{Val: 4}
	t33 := TreeNode{Val: 2}
	// t34 := TreeNode{Val: 3}

	t11.Left = &t21
	t11.Right = &t22
	t21.Left = &t31
	// t21.Right = &t32
	t22.Left = &t33
	// t22.Right = &t34
	print(&t11)
	fmt.Printf("%t\n", isSymmetric(&t11))
}