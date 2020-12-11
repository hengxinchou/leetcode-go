/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

package main

import "fmt"

type Node struct {
    Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	
	// queue := []*Node{}
	// queue = append(queue, root)
	// for len(queue) > 0 {
	// 	size := len(queue)
	// 	var pre *Node = nil
	// 	for i := 0 ; i < size; i ++ {
	// 		cur := queue[0]
	// 		queue = queue[1:len(queue)]
	// 		if pre != nil {
	// 			pre.Next = cur
	// 		}
	// 		pre = cur
	// 		if cur.Left != nil {
	// 			queue = append(queue, cur.Left)
	// 		}
	// 		if cur.Right != nil {
	// 			queue = append(queue, cur.Right)
	// 		}
	// 	}
	// }
	// return root

	// 学习网上的
	cur := root
	for cur != nil {
		//遍历当前层的时候，为了方便操作在下一
		//层前面添加一个哑结点（注意这里是访问
		//当前层的节点，然后把下一层的节点串起来）
		dummy := &Node{}
		//pre表示访下一层节点的前一个节点
		pre := dummy
		//然后开始遍历当前层的链表
		for cur != nil {
			if cur.Left != nil {
				//如果当前节点的左子节点不为空，就让pre节点
				//的next指向他，也就是把它串起来
				pre.Next = cur.Left
				//然后再更新pre
				pre = pre.Next
			}
			//同理参照左子树
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			//继续访问这一行的下一个节点
			cur = cur.Next
		}
		//把下一层串联成一个链表之后，让他赋值给cur，
		//后续继续循环，直到cur为空为止
		cur = dummy.Next
	}
	return root;

}
// @lc code=end

func printQueue(queue []*Node){
	fmt.Printf("print queue:  ")
   for i := range queue {
	   fmt.Printf("%d ", queue[i].Val)
   }
   fmt.Println()
}

func print(root *Node){
	fmt.Printf("%d ,", root.Val)
	if root.Next != nil {
		fmt.Printf("next is %d ", root.Next.Val)
	} else {
		fmt.Printf("next is null")
	}
	fmt.Println()
	if root.Left != nil {
	  print(root.Left)
    }
    if root.Right != nil {
	  print(root.Right)
	}
}


func main (){
	l11 := Node{Val: 1}
	l21 := Node{Val: 2} 
	l22 := Node{Val: 3} 
	l31 := Node{Val: 4} 
	l32 := Node{Val: 5}
	l33 := Node{Val: 7}

	l11.Left = &l21
	l11.Right = &l22
	l21.Left = &l31
	l21.Right = &l32
	l22.Right = &l33
	print(&l11)
	connect(&l11)
	fmt.Println()
	print(&l11)
	fmt.Println()
}
