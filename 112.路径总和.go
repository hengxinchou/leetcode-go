/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

func hasPathSum(root *TreeNode, sum int) bool {
	// return true
	if root == nil  {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return  root.Val == sum 
	}	
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
// @lc code=end
func print(root *TreeNode){
	// fmt.Printf("%d ", root.Val)
	// if root.Left != nil {
	// 	print(root.Left)
	// }
	// if root.Right != nil {
	// 	print(root.Right)
	// }
	
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

func main(){
	t11 := TreeNode{Val: 5}
	t21 := TreeNode{Val: 4}
	t22 := TreeNode{Val: 8}
	t31 := TreeNode{Val: 11}
	// t32 := TreeNode{Val: 13}
	t33 := TreeNode{Val: 13}
	t34 := TreeNode{Val: 4}
	t41 := TreeNode{Val: 7}
	t42 := TreeNode{Val: 2}
	t43 := TreeNode{Val: 1}

	t11.Left = &t21
	t11.Right = &t22
	t21.Left = &t31
	// t21.Right = &t32
	t22.Left = &t33
	t22.Right = &t34

	t31.Left = &t41
	t31.Right = &t42

	t34.Right = &t43
	print(&t11)
	fmt.Println(hasPathSum(&t11, 22))


	fmt.Println(hasPathSum(&t11, 18))
	fmt.Println(hasPathSum(&t11, 26))
	fmt.Println(hasPathSum(&t11, 25))
}