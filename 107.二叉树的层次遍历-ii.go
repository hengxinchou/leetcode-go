/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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


func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	// stack = append(stack, root)
	res := [][]int{}

	for len(stack) > 0 {
		size := len(stack)
		tmpArray := []int{}
		for i := 0 ; i < size ; i++ {
			tmp := stack[0]
			tmpArray = append(tmpArray, tmp.Val)
			stack = stack[1:]
			// fmt.Printf("%d ", tmp.Val)
			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}
		}
		res = append([][]int{tmpArray}, res...)
		// fmt.Println()
	}
	return res
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
	t11 := TreeNode{Val: 3}
	t21 := TreeNode{Val: 9}
	t22 := TreeNode{Val: 20}
	t33 := TreeNode{Val: 15}
	t34 := TreeNode{Val: 7}

	t11.Left = &t21
	t11.Right = &t22
	t22.Left = &t33
	t22.Right = &t34
	print(&t11)
	fmt.Printf("%v\n", levelOrderBottom(&t11))
}

