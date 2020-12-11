/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func print(root *TreeNode) {
	stack := []*TreeNode{root}

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

func main() {
	l11 := TreeNode{Val: 1}
	l21 := TreeNode{Val: 2} 
	l22 := TreeNode{Val: 3} 
	// l31 := TreeNode{Val: 5} 
	// l32 := TreeNode{Val: 1}

	l11.Left = &l21
	l11.Right = &l22
	// l21.Left = &l31
	// l21.Right = &l32

	print(&l11)
	fmt.Printf("%d\n", sumNumbers(&l11))
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	a := getNumbers2(root) 
	sum := 0
	for _, v := range a {
		// fmt.Printf("v is %s\n", v)
		tmp, _ := strconv.Atoi(v)
		sum += tmp
	}
	return sum
}


func getNumbers2(root *TreeNode) []string {
	if root.Left == nil && root.Right == nil {
		a :=strconv.Itoa(root.Val)
		return []string{a}
	}

	res := []string{}
	
	// tmp := root.Val * 10
	// fmt.Printf("root is %d, tmp is %d\n", root.Val, tmp)
	tmp := strconv.Itoa(root.Val)
	if root.Left != nil {
		tmp2 := getNumbers2(root.Left)
		// fmt.Printf("Left tmp2 is %v\n", tmp2)
		for _, v := range tmp2 {
			res = append(res,  tmp + v)
		}
	}
	if root.Right != nil {
		tmp2 := getNumbers2(root.Right)
		// fmt.Printf("Right tmp2 is %v\n", tmp2)
		for _, v := range tmp2 {
			res = append(res,  tmp + v)
		}
	}
	return res
}
// @lc code=end

// func getNumbers(stack []*TreeNode) (res []int) {
// 	for _, v := range stack {
// 		res = append(res, v.Val)
// 	}
// 	return
// }

