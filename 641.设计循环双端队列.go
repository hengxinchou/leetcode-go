/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
package main

import "fmt"

type MyCircularDeque struct {
	head int
	tail int
	size int
    data []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{ 
		head: 0,
		tail: 0, 
		size: k+1, 
		data: make([]int, k+1, k+1),
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	my_head := (this.head + this.size - 1) % this.size
	this.data[my_head] = value
	this.head = my_head
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// fmt.Printf("func InsertLast, head is %d, tail is %d\n", this.head, this.tail)
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail  + 1) % this.size 
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head  + 1) % this.size 

	return true

}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail + this.size - 1) % this.size
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[(this.tail + this.size -1) % this.size]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail + 1 ) % this.size == this.head
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

func main (){
//   k := 3
//   circularQueue := Constructor(k);
//   fmt.Println("null")
//   fmt.Printf("InsertLast is %t\n", circularQueue.InsertLast(1))
//   fmt.Printf("InsertLast is %t\n", circularQueue.InsertLast(2))
//   fmt.Printf("InsertFront is %t\n", circularQueue.InsertFront(3))
//   fmt.Printf("InsertFront is %t\n", circularQueue.InsertFront(4))
//   fmt.Printf("GetRear is %d\n", circularQueue.GetRear())
//   fmt.Printf("IsFull is %t\n", circularQueue.IsFull())
//   fmt.Printf("DeletaLast is %t\n", circularQueue.DeleteLast())
//   fmt.Printf("InsertFront is %t\n", circularQueue.InsertFront(4))
//   fmt.Printf("GetFront is %d\n", circularQueue.GetFront())

  k := 8
  circularQueue := Constructor(k);
  fmt.Println("null")
  fmt.Printf("InsertFront is %t\n", circularQueue.InsertFront(5))
  fmt.Printf("GetFront is %d\n", circularQueue.GetFront())
  fmt.Printf("isEmpty is %t\n", circularQueue.IsEmpty())
  fmt.Printf("DeleteFront is %t\n", circularQueue.DeleteFront())
  fmt.Printf("InsertLast is %t\n", circularQueue.InsertLast(3))
  fmt.Printf("GetRear is %d\n", circularQueue.GetRear())
  fmt.Printf("InsertLast is %t\n", circularQueue.InsertLast(7))
  fmt.Printf("InsertFront is %t\n", circularQueue.InsertFront(7))
  fmt.Printf("DeletaLast is %t\n", circularQueue.DeleteLast())
  fmt.Printf("InsertLast is %t\n", circularQueue.InsertLast(4))
  fmt.Printf("isEmpty is %t\n", circularQueue.IsEmpty())
}