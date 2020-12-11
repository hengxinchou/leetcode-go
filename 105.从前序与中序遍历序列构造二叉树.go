/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
	// var a []int	
	a1 := []int{3,9,20,15,7}
	a2 := []int{9,3,15,20,7}
	a3 := buildTree(a1, a2)
	print(a3)
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1  {
		return nil
	}

	rootValue := preorder[0]

	var i int
	for i = 0 ; i < len(inorder) ; i++ {
		if inorder[i] == rootValue  {
			break
		}
	}
	// fmt.Printf("%v, %v, %d\n", preorder,inorder,i)
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
// @lc code=end

