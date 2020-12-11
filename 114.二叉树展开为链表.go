/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
	l11 := TreeNode{Val: 1}
	l21 := TreeNode{Val: 2} 
	l22 := TreeNode{Val: 5} 
	l31 := TreeNode{Val: 3} 
	l32 := TreeNode{Val: 4}
	l33 := TreeNode{Val: 6}

	l11.Left = &l21
	l11.Right = &l22
	l21.Left = &l31
	l21.Right = &l32
	l22.Right = &l33
	print(&l11)
	flatten(&l11)
	print(&l11)
	// fmt.Printf("%v\n", res)	

}

func print(root *TreeNode) {	
	queue := []*TreeNode{}
	queue = append(queue, root)
	
	for len(queue) > 0 {
		size := len(queue)
		for i := 0 ; i < size; i ++ {	
			cur := queue[0]
			fmt.Printf("%d ", cur.Val)
			queue = queue[1:len(queue)]
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

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}	
	
	stack := []*TreeNode{}
	cur := root
    queue := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)	
			queue = append(queue, cur)
			cur = cur.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1 ]
		cur = tmp.Right
	}
	prev := root
	// for i := 0 ; i < len(queue) ; i++ {
	// 	fmt.Printf("%d ", queue[i].Val)
	// }
	// fmt.Println()
	for i := 1 ; i < len(queue) ; i++ {
		prev.Left = nil
		prev.Right = queue[i]
		prev = queue[i]
	}
	return
}
// @lc code=end

