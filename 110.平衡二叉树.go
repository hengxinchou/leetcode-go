/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
 

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}


func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+height(root.Left), 1+height(root.Right))
}

func isBalanced2(root *TreeNode) bool {
	tmp := height(root.Left) - height(root.Right) 
	if tmp <= 1 && tmp >= -1 {
		return true
	} else {
		return false
	}
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		fmt.Printf("root is %d\n", root.Val)
		if !isBalanced2(root) {
			return false
		}
		stack = stack[0:len(stack)-1]
		root = root.Right
	}

	return true
}
// @lc code=end

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
	// t21 := TreeNode{Val: 1}
	t22 := TreeNode{Val: 7}
	// t33 := TreeNode{Val: 6}
	// t34 := TreeNode{Val: 9}
	// t44 := TreeNode{Val: 9}

	// t11.Left = &t21
	t11.Right = &t22
	// t22.Left = &t33
	// t22.Right = &t34
	// t34.Right = &t44
	print(&t11)
	fmt.Println(isBalanced(&t11))
}
