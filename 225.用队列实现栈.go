/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	data []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{data: []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.data = append(this.data, x)
	return
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	size := len(this.data)
	x := this.data[size]
	this.data = this.data[0:size-1]
	return
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0 
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

