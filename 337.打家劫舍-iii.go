/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := []int{}

	stack := []*TreeNode{}
	stack = append(stack, root)
	// flag := true
	for len(stack) > 0 {
		size := len(stack)

		sum := 0
		for i := 0 ; i < size ; i++ {
			tmp := stack[0]
			stack = stack[1:]
			// fmt.Printf("%d ", tmp.Val)
			sum += tmp.Val
			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}
		}
		// fmt.Println()
		res = append(res, sum)
	}
	
	return myRob(res)
}
// @lc code=end

func myRob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		max(nums[0], nums[1])
	}
	// fmt.Printf("nums are %v\n", nums)
	a := make([]int, len(nums))
	a[0] = nums[0]
	a[1] = max(nums[1], nums[0])
	for i := 2 ; i < len(nums) ; i++ {
		a[i] = max(a[i-2] + nums[i], a[i-1])
	}
	return a[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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
	t11 := TreeNode{Val: 2}
	t21 := TreeNode{Val: 1}
	t22 := TreeNode{Val: 3}
	t31 := TreeNode{Val: 4}
	// t34 := TreeNode{Val: 9}

	t11.Left = &t21
	t11.Right = &t22
	t21.Right = &t31
	// t22.Right = &t34
	print(&t11)
	fmt.Println(rob(&t11))
}

