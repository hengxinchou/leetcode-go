/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用递归
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}

	if root.Left != nil {
		tmp := preorderTraversal(root.Left)
		res = append(res, tmp...)
	}
	if root.Right != nil {
		tmp := preorderTraversal(root.Right)
		res = append(res, tmp...)
	}
	return res
}

// 使用栈
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := []int{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			fmt.Printf("cur is %d\n", cur.Val)
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		tmp := stack[len(stack)-1]
		fmt.Printf("tmp is %d\n", tmp.Val)
		stack = stack[:len(stack)-1]
		cur = tmp.Right
	}
	return res
}

// @lc code=end
func main() {
	l11 := TreeNode{Val: 6}
	l21 := TreeNode{Val: 2}
	l22 := TreeNode{Val: 8}
	l31 := TreeNode{Val: 0}
	l32 := TreeNode{Val: 4}
	l33 := TreeNode{Val: 7}
	l34 := TreeNode{Val: 9}
	l41 := TreeNode{Val: 3}
	l42 := TreeNode{Val: 5}

	l11.Left = &l21
	l11.Right = &l22
	l21.Left = &l31
	l21.Right = &l32
	l22.Left = &l33
	l22.Right = &l34
	l32.Left = &l41
	l32.Right = &l42

	a := preorderTraversal(&l11)
	fmt.Printf("%v\n", a)
}
