/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
 

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
	if root.Left != nil {
		return minDepth(root.Left) + 1
	} else {
		return minDepth(root.Right) + 1
	}
}
// @lc code=end

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func print(root *TreeNode){
	if root == nil {
		return 
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		size := len(stack)
		for i := 0 ; i < size ; i++ {
			tmp := stack[0]
			stack = stack[1:]
			fmt.Printf("%d ", tmp.Val)
			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}
		}
		fmt.Println()
	}
}


// @lc code=end
func main(){
	t11 := TreeNode{Val: 5}
	t21 := TreeNode{Val: 1}
	// t22 := TreeNode{Val: 7}
	t33 := TreeNode{Val: 6}
	// t34 := TreeNode{Val: 9}
	t44 := TreeNode{Val: 9}

	// t11.Left = &t21
	// t11.Right = &t22
	// t22.Left = &t33
	// t22.Right = &t34
	// t34.Right = &t44
	t11.Right = &t21
	t21.Right = &t33
	t33.Right = &t44
	print(&t11)
	fmt.Printf("min height is %d\n", minDepth(&t11))
}
