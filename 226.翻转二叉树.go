/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	return root

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
func main(){
	t11 := TreeNode{Val: 4}
	t21 := TreeNode{Val: 2}
	t22 := TreeNode{Val: 7}
	t31 := TreeNode{Val: 1}
	t32 := TreeNode{Val: 3}
	t33 := TreeNode{Val: 6}
	t34 := TreeNode{Val: 9}

	t11.Left = &t21
	t11.Right = &t22
	t21.Left = &t31
	t21.Right = &t32
	t22.Left = &t33
	t22.Right = &t34

	print(&t11)
	fmt.Println()

	a := invertTree(&t11)
	print(a)
	fmt.Println()
}
