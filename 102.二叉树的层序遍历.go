/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	
	res := [][]int{}

	queue := []*TreeNode{}
	queue = append(queue, root)
	
	for len(queue) > 0 {
		printQueue(queue)
		size := len(queue)
		a := []int{}
		for i := 0 ; i < size; i ++ {	
			cur := queue[0]
			queue = queue[1:len(queue)]
			a = append(a, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, a)
	}
	return res
}
// @lc code=end
func print(root *TreeNode){
	fmt.Printf("%d ", root.Val)

	fmt.Println()
	if root.Left != nil {
	  print(root.Left)
    }
    if root.Right != nil {
	  print(root.Right)
	}
}

func printQueue(queue []*TreeNode){
	fmt.Printf("print queue:  ")
   for i := range queue {
	   fmt.Printf("%d ", queue[i].Val)
   }
   fmt.Println()
}

func main (){
	l11 := TreeNode{Val: 1}
	l21 := TreeNode{Val: 2} 
	l22 := TreeNode{Val: 3} 
	l31 := TreeNode{Val: 4} 
	l32 := TreeNode{Val: 5}
	l33 := TreeNode{Val: 7}

	l11.Left = &l21
	l11.Right = &l22
	l21.Left = &l31
	l21.Right = &l32
	l22.Right = &l33
	res := levelOrder(&l11)
	fmt.Printf("%v\n", res)
}

