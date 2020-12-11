/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func main() {
	t11 := &TreeNode{Val: 5}	
	t21 := &TreeNode{Val: 4}	
	t22 := &TreeNode{Val: 8}	
	t31 := &TreeNode{Val: 11}	
	t32 := &TreeNode{Val: 13}	
	t33 := &TreeNode{Val: 4}	
	t41 := &TreeNode{Val: 7}	
	t42 := &TreeNode{Val: 2}	
	t43 := &TreeNode{Val: 5}	
	t44 := &TreeNode{Val: 1}	

	t11.Left = t21
	t11.Right = t22

	t21.Left = t31
	t22.Left = t32
	t22.Right = t33

	t31.Left = t41
	t31.Right = t42
	t33.Left = t43
	t33.Right = t44
	print(t11)

	var res [][]int
	res = pathSum(t11, 22)
	fmt.Println(res)

	res = pathSum(nil, 22)
	fmt.Println(res)

	t111 := &TreeNode{Val: -2}
	t112 := &TreeNode{Val: -3}
	t111.Right = t112
	print(t111)
	res = pathSum(t111, -5)
	fmt.Println(res)
}

func print(root *TreeNode){

	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0 ;i < size ; i ++ {
			tmp := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", tmp.Val)

			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}

			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		fmt.Println()
	}
}


func pathSum1(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}	

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{[]int{root.Val}}
		} else  {
			return nil
		}
	}

	left  := sum - root.Val
	res := [][]int{}
	if root.Left != nil {
		tmp := pathSum(root.Left, left)
		if tmp != nil {
			for _, v := range tmp {
				tmp1 := append([]int{root.Val}, v...)
				res = append(res, tmp1)
			}
		}
	}
	if root.Right != nil {
		tmp := pathSum(root.Right, left)
		if tmp  != nil {
			for _, v := range tmp {
				tmp1 := append([]int{root.Val}, v...)
				res = append(res, tmp1)
			}
		}
	}
	return res
}

func pathSum2(root *TreeNode, sum int) [][]int {	
	if root == nil {
		return [][]int{}
	}

	var dfs func(*TreeNode, int, []int)
	res := [][]int{}
	dfs = func (tree *TreeNode, acculate int, accs []int) {
		acculate += tree.Val
		accs = append(accs, tree.Val)
		// fmt.Printf("tree is %d, %d\n", tree.Val, acculate)
		if sum == acculate {
			if tree.Left == nil && tree.Right == nil {
				res = append(res, accs)
				return
			 } 
		}	


		if tree.Left != nil {
			tmp := append([]int{}, accs...)
			dfs(tree.Left, acculate, tmp)
		}
		if tree.Right != nil {
			tmp := append([]int{}, accs...)
			dfs(tree.Right, acculate, tmp)
		}
	}
	dfs(root, 0, []int{})
	// fmt.Printf("count is %d\n", count)
	return res
}
// @lc code=end

func pathSum(root *TreeNode, sum int) (ans [][]int) {
    path := []int{}
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, left int) {
        if node == nil {
            return
        }
        left -= node.Val
		path = append(path, node.Val)
		// niubi!!!!
        defer func() { 
			// fmt.Printf("before path is %v\n", path)
			path = path[:len(path)-1] 
			// fmt.Printf("after path is %v\n", path)
		}()
        if node.Left == nil && node.Right == nil && left == 0 {
            ans = append(ans, append([]int(nil), path...))
            return
        }
        dfs(node.Left, left)
        dfs(node.Right, left)
    }
    dfs(root, sum)
    return
}