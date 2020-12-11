/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
 
 // 不能用前序和中序的方法，因为会重复访问同一个节点。
 // 因为总是访问左，然后访问根；或者先访问根节点，再访问左节点，
 // 但后序就不可以，不能先访问左节点，再访问右节点，然后访问根节点，根节点还得判断是否有左右节点
 // 这样就陷入死循环了。
 //
 //  可以反过来想，因为 后序是 左，右， 根节点
 // 可以先访问 根， 右，左， 然后反向输出，也是跟结果一样的
 func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	res1 := []int{}

	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res1 = append(res1, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		
		tmp := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]
		cur = tmp.Left
	}

	res2 := make([]int, len(res1))
	for i := len(res1) ; i > 0 ; i-- {
		res2[len(res1) - i] = res1[i-1]
	}
	return res2
}
// @lc code=end

func print(a []*TreeNode) {
	fmt.Printf("stack is ")
	for _, b := range a {
		fmt.Printf("%d ", b.Val)
	}
	fmt.Println()
}

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

	a := postorderTraversal(&l11)
	fmt.Printf("%v\n", a)
}
