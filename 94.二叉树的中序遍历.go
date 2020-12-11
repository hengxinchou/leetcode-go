/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

func inorderTraversal(root *TreeNode) []int {
	// a := []int{}
	// if root.Left != nil {
	// 	b := inorderTraversal(root.Left)
	// 	a  = append(a, b...)
	// }
	// a = append(a, root.Val)
	// if root.Right != nil {
	// 	c := inorderTraversal(root.Right)
	// 	a = append(a, c...)
	// }
	// return a
	stack := make([]*TreeNode, 0)
	res := []int{}
	cur := root
	for cur != nil || len(stack) > 0 {
		// fmt.Printf("first for,  %d \n", cur.Val)
		// left := cur.Left
		for  cur != nil {
			// fmt.Printf("second for, %d \n", cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		// fmt.Println("to get tmp")
		tmp := stack[len(stack)-1]
		// fmt.Println("to get tmp")
		stack = stack[:len(stack)-1]
		res = append(res, tmp.Val)
		cur = tmp.Right

		// fmt.Printf("res is %v\n", res)
	} 
	return res
}
// @lc code=end

func main(){
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

	a := inorderTraversal(&l11)
	fmt.Printf("%v\n", a)
}
