/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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


func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	size := len(nums)
	mid := size / 2
	low := 0
	upper := size
	
	// fmt.Printf("nums are %v, mid is %d, low is %d, upper is %d\n", nums, mid, low, upper)
	root := &TreeNode{Val: nums[mid]}

	if low < mid {
		tmp := sortedArrayToBST(nums[low: mid])
		root.Left = tmp
	} 

	if upper > mid {
		tmp := sortedArrayToBST(nums[mid+1: upper])
		root.Right = tmp
	}
	return root
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


func main(){
  a := []int{-10,-3,0,5,9}
  b := sortedArrayToBST(a)
  print(b)
}