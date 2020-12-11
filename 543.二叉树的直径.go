/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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


func main() {
	t11 := &TreeNode{Val:1}			
	t21 := &TreeNode{Val:2}			
	t22 := &TreeNode{Val:3}			
	t31 := &TreeNode{Val:4}			
	t32 := &TreeNode{Val:5}			

	t11.Left = t21
	t11.Right = t22
	t21.Left = t31
	t21.Right = t32

	print(t11)
	fmt.Println(diameterOfBinaryTree(t11))
}

func print(root *TreeNode) {
	// if root == nil {
	// 	return 
	// }
	// stack := []*TreeNode{}
	// stack = append(stack, root)
	// for len(stack) > 0 {
	// 	size := len(stack)
	// 	for i := 0 ; i < size ; i++ {
	// 		tmp := stack[0]
	// 		stack = stack[1:]
	// 		fmt.Printf("%d ", tmp.Val)
	// 		if tmp.Left != nil {
	// 			stack = append(stack, tmp.Left)
	// 		}
	// 		if tmp.Right != nil {
	// 			stack = append(stack, tmp.Right)
	// 		}
	// 	}
	// 	fmt.Println()
	// }


	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)

		for i := 0 ; i < size ; i++ {
			tmp := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", tmp.Val)
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func diameterOfBinaryTree(root *TreeNode) int {	
	if root == nil {
		return 0
	}

	maxV := 0

	queue := []*TreeNode{root}
	if len(queue) > 0 {
		size := len(queue)

		for i := 0 ; i < size ; i ++ {
			tmp := queue[0]
			queue = queue[1:]
			tmp2 := height(root.Left)  +  height(root.Right) + 1
			if tmp2 > maxV {
				maxV = tmp2
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
	}

	return maxV -1 
}
// @lc code=end
 
func height(root  *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}