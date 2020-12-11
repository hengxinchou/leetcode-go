/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
package main
import "fmt"

func main() {

	a := Constructor(2)
	a.Put(1,1)
	// a.Print()
	a.Put(2,2)
	a.Print()
	b := a.Get(1)
	fmt.Printf("get 1 is %d\n", b)
	a.Put(3,3)
	a.Print()
	a.Put(3,6)
	a.Print()
	c := a.Get(1)
	fmt.Printf("get 1 is %d\n", c)
	a.Print()
	
}
type Item struct {
	key int
	value int
}

type LRUCache struct {
	capacity int
	data []Item
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, data: []Item{}}
}


func (this *LRUCache) Get(key int) int {
	for i, item := range this.data {
		if item.key == key {
			this.data = append(append([]Item{item}, this.data[0:i]...), this.data[i+1:]...)
			return item.value
		}
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	item := Item{key: key, value: value}
	for i, iter := range this.data {
		if iter.key == key {
			// item = Item{key:key, value: value}
			this.data = append(append([]Item{item}, this.data[0:i]...), this.data[i+1:]...)
			return
		}
	}
	if len(this.data) < this.capacity {
		this.data = append([]Item{item}, this.data...)
	} else {
		this.data = append([]Item{item}, this.data[0:len(this.data)-1]...)
	}
}

func (this *LRUCache) Print() {
	fmt.Printf("capacity: %d\n", this.capacity)
	for i := 0 ; i < len(this.data) ; i++ {
		fmt.Printf("%d, %d\n", this.data[i].key, this.data[i].value)
	}
	fmt.Println("finish")
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end




