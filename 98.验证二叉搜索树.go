/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

func getChild(root *TreeNode) []int{
	if root == nil {
		return []int{}
	}
	res := []int{}
	if root.Left != nil {
		tmp := getChild(root.Left)
		res = append(res, tmp...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		tmp := getChild(root.Right)
		res = append(res, tmp...)
	}
	return res
}



func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	tmp := getChild(root)
	for i := 0 ; i < len(tmp) ; i ++ {
		if i > 0 && tmp[i] <= tmp[i-1] {
			return false
		}
	}
	return true
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
	t22 := TreeNode{Val: 7}
	t33 := TreeNode{Val: 6}
	t34 := TreeNode{Val: 9}

	t11.Left = &t21
	t11.Right = &t22
	t22.Left = &t33
	t22.Right = &t34
	print(&t11)
	fmt.Println(isValidBST(&t11))
}
