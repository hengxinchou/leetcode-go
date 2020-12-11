/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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

func print(root *TreeNode) {
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
	}
	fmt.Println()
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
// @lc code=end

func main(){
	t11 := TreeNode{Val: 1}
	t21 := TreeNode{Val: 2}
	t22 := TreeNode{Val: 2}
	t31 := TreeNode{Val: 2}
	t32 := TreeNode{Val: 4}
	t33 := TreeNode{Val: 4}
	t34 := TreeNode{Val: 2}

	t11.Left = &t21
	t11.Right = &t22
	t21.Left = &t31
	t21.Right = &t32
	t22.Left = &t33
	t22.Right = &t34
	print(&t11)
	fmt.Printf("%t\n", isSameTree(&t11, &t11))
	fmt.Printf("%t\n", isSameTree(&t11, &t21))
}
