/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
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

func isAncestor(a, b *TreeNode) bool {
	if a == nil {
		return false 
	}

	if a == b {
		return true
	}

	if a.Left == b {
		return true
	}
	if a.Right == b {
		return true
	}

	if a.Left != nil && isAncestor(a.Left, b) {
		return true
	} 
	if a.Right != nil && isAncestor(a.Right, b){
		return true
	}
	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	fmt.Printf("p is %d, q is %d\n", p.Val, q.Val)
	cur := root
	// prev := root
	for cur != nil {
		fmt.Printf("cur is %d\n", cur.Val)
		left := cur.Left
		if isAncestor(left, p) && isAncestor(left, q) {
			fmt.Printf("left is %d\n",left.Val)
			// prev = cur
			cur = left
			continue
		} 
		right := cur.Right
		if isAncestor(right, p) && isAncestor(right, q) {
			fmt.Printf("right is %d\n", right.Val)
			// prev = cur
			cur = right
			continue
		}
		break
	}
	return cur
}
// @lc code=end

func print(root *TreeNode){
	fmt.Printf("%d ", root.Val)
	if root.Left != nil {
		print(root.Left)
	}
	
	if root.Right != nil {
		print(root.Right)
	}
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
	print(&l11)
	fmt.Println()

	c := lowestCommonAncestor(&l11, &l21, &l32)
	fmt.Printf("result is %d\n", c.Val)
}
