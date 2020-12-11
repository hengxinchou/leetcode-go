/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth1, depth2 := 0, 0
	if root.Left != nil {
		depth1 = maxDepth(root.Left)
	}
	if root.Right != nil {
		depth2 = maxDepth(root.Right)
	}

	if depth1 > depth2 {
		return 1 + depth1
	} else {
		return 1 + depth2
	}
}


func print(root *TreeNode){
	fmt.Printf("%d ", root.Val)
	if root.Left != nil {
	  print(root.Left)
    }
    if root.Right != nil {
	  print(root.Right)
	}
}

// @lc code=end
func main (){
	l11 := TreeNode{Val: 3}
	l21 := TreeNode{Val: 9} 
	l22 := TreeNode{Val: 20} 
	l31 := TreeNode{Val: 15} 
	l32 := TreeNode{Val: 7}
	l41 := TreeNode{Val: 8}

	l11.Left = &l21
	l11.Right = &l22
	l22.Left = &l31
	l22.Right = &l32
	l32.Left = &l41

	print(&l11)
	fmt.Println()
	fmt.Printf("max depth is %d\n", maxDepth(&l11))
}
