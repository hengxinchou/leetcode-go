/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
package main
import "fmt"
type MyCircularQueue struct {
	Head int
	Tail int 
	Size int
	Array []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{}
	queue.Head = 0
	queue.Tail = 0
	queue.Size = k+1
	queue.Array = make([]int, k + 1, k + 1)
	return queue
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
// 因为刚开始也是 tail 等于head，如果tail不间隔一个,则无法区分开始的情况；因此浪费一个空间
// 必须是 先 + 1 ，再取余；不能先取余，再加1，因为这样就没办法进位了。
func (this *MyCircularQueue) EnQueue(value int) bool {
	if (this.Tail +1) % this.Size == this.Head {
		return false
	}
	this.Array[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Size 
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty()  {
		return false
	}
	this.Head = (this.Head + 1) % this.Size
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Array[this.Head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	 return this.Array[(this.Tail + this.Size - 1 ) % this.Size]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail 
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.Tail+ 1) % this.Size == this.Head
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

func main (){
  k := 6
  circularQueue := Constructor(k);
  fmt.Println("null")
  fmt.Printf("Enqueue is %t\n", circularQueue.EnQueue(6))
  fmt.Printf("rear is %d\n", circularQueue.Rear())
  fmt.Printf("rear is %d\n", circularQueue.Rear())
  fmt.Printf("Dequeue is %t\n", circularQueue.DeQueue())
  fmt.Printf("Enqueue is %t\n", circularQueue.EnQueue(5))
  fmt.Printf("Rear is %d\n", circularQueue.Rear())
  fmt.Printf("DeQueue is %t\n", circularQueue.DeQueue())
  fmt.Printf("DeQueue is %d\n", circularQueue.Front())
  fmt.Printf("DeQueue is %t\n", circularQueue.DeQueue())
  fmt.Printf("DeQueue is %t\n", circularQueue.DeQueue())
  fmt.Printf("DeQueue is %t\n", circularQueue.DeQueue())
}