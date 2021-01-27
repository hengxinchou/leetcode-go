/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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

func main() {
	l11 := TreeNode{Val: 3}
	l21 := TreeNode{Val: 9}
	l22 := TreeNode{Val: 20}
	l31 := TreeNode{Val: 15}
	l32 := TreeNode{Val: 7}

	l11.Left = &l21
	l11.Right = &l22
	l22.Left = &l31
	l22.Right = &l32
	print(&l11)
	a := zigzagLevelOrder(&l11)
	fmt.Println(a)
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	flag := true
	for len(stack) > 0 {
		size := len(stack)
		tmp := []int{}
		for i := 0; i < size; i++ {
			cur := stack[0]
			if flag {
				tmp = append(tmp, cur.Val)
			} else {
				tmp = append([]int{cur.Val}, tmp...)
			}
			stack = stack[1:]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		flag = !flag
		res = append(res, tmp)
	}
	return res
}

func print(root *TreeNode) {
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			fmt.Printf("%d ", cur.Val)
			queue = queue[1:len(queue)] // 弹出对头元素
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		fmt.Println()
	}
}

// @lc code=end
