/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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



type BSTIterator struct {
	Root *TreeNode
	Cur int
	Data []int
}


func Constructor(root *TreeNode) BSTIterator {
	data := InOrder(root)
	return BSTIterator{Root: root, Data: data, Cur: -1}
}

func InOrder(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		tmp := stack[len(stack) - 1]
		res = append(res, tmp.Val)
		stack = stack[0:len(stack)-1]		
		cur = tmp.Right
		// if cur != nil {
		// 	res = append(res, cur.Val)
		// }
	}
	return res
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.Cur++
	return this.Data[this.Cur]
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return (this.Cur + 1 ) < len(this.Data)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end


func main() {
	t11 := &TreeNode{Val:7}
	t21 := &TreeNode{Val:3}
	t22 := &TreeNode{Val:15}
	// t31 := &TreeNode{Val:1}
	// t32 := &TreeNode{Val:4}
	t33 := &TreeNode{Val:9}
	t34 := &TreeNode{Val:20}
	t11.Left = t21
	t11.Right = t22
	// t21.Left = t31
	// t21.Right = t32
	t22.Left = t33
	t22.Right = t34

	b1 := Constructor(t11)

	fmt.Println(b1.Data)
	fmt.Println(b1.Next())
	fmt.Println(b1.Next())
	fmt.Println(b1.HasNext())
	fmt.Println(b1.Next())
	fmt.Println(b1.HasNext())
	fmt.Println(b1.Next())
	fmt.Println(b1.HasNext())
	fmt.Println(b1.Next())
	fmt.Println(b1.HasNext())
}