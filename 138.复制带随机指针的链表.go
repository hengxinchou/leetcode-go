/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
package main

import "fmt"

type Node struct {
    Val int
    Next *Node
	Random *Node
}

func main() {
	v1 := &Node{Val: 7}
	v2 := &Node{Val: 13}
	v3 := &Node{Val: 11}
	v4 := &Node{Val: 10}
	v5 := &Node{Val: 1}

	v1.Next = v2
	v2.Next = v3
	v3.Next = v4
	v4.Next = v5

	v1.Random = nil
	v2.Random = v1
	v3.Random = v5
	v4.Random = v3
	v5.Random = v1

	print(v1)

	var vnew *Node
	vnew = copyRandomList(v1)
	print(vnew)

}

func print(head *Node) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
	
	cur = head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		tmp := cur.Random
		if tmp != nil {
		    fmt.Printf("=> %d ", tmp.Val)
		} else {
			fmt.Printf("=> nil ")
		}
		fmt.Printf(", ")
		cur = cur.Next
	}
	fmt.Println()
}


func copyRandomList(head *Node) *Node {
   if head == nil {
	   return head
   }

   headB := &Node{}
   prev := headB

   a := []*Node{}
   b := []*Node{}

   cur := head

   for cur != nil {
		a = append(a, cur)
		
		copyNode := &Node{Val: cur.Val}
		prev.Next = copyNode
		b = append(b, copyNode)
		prev = copyNode

		cur = cur.Next		
   }

   cur = head
	for i := 0 ;i < len(a); i ++ {
		if a[i].Random == nil {
			b[i].Random = nil
			continue
		}

		for j := 0 ; j < len(a) ;j ++ {
			if a[i].Random == a[j] {
				b[i].Random = b[j]
				break
			}
		}
	}

	return headB.Next
}
// @lc code=end

