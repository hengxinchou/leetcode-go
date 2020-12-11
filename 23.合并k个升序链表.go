/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main
import "fmt"
import "container/heap"

type ListNode struct {
    Val int
	Next *ListNode
}

type heapList []*ListNode  

func (h heapList) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h heapList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapList) Len() int {
	return len(h)
}

func (h *heapList) Pop() (v interface{}) {
	v = (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func (h *heapList) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}



func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	myHeapList := heapList{}
	for i := 0 ; i < len(lists) ; i++ {
		if lists[i] != nil {
			myHeapList = append(myHeapList, lists[i])
		}
	}
	heap.Init(&myHeapList)
	
	root := &ListNode{}
	cur := root
	for myHeapList.Len() != 0 {
		tmp := heap.Pop(&myHeapList).(*ListNode)
		// fmt.Printf("small is %d\n", tmp.Val)
		cur.Next = tmp
		cur = tmp
		wait := tmp.Next
		if wait != nil {
			heap.Push(&myHeapList, wait)
		} else {
			continue
		}
	}
	return root.Next
}
// @lc code=end

func print(l *ListNode){
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main(){
	a11 := ListNode{Val: 1}
	a12 := ListNode{Val: 4}
	a13 := ListNode{Val: 5}
	a11.Next = &a12
	a12.Next = &a13
	print(&a11)

	a21 := ListNode{Val: 1}
	a22 := ListNode{Val: 3}
	a23 := ListNode{Val: 4}
	a21.Next = &a22
	a22.Next = &a23
	print(&a21)

	a31 := ListNode{Val: 2}
	a32 := ListNode{Val: 6}
	a31.Next = &a32
	print(&a31)

	// res := mergeKLists([]*ListNode{&a11, &a21, &a31})
	// print(res)

	// res2 := mergeKLists(nil)
	// print(res2)
	res3 := mergeKLists([]*ListNode{&a11, nil})
	print(res3)
}

