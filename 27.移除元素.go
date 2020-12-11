/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
package main

import "fmt"

func removeElement(nums []int, val int) int {
	size := len(nums)
	for  i := 0 ; i < size ; {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			fmt.Printf("append is %v\n", nums)
			size--
		} else {
			i++
		}
	}
	return len(nums)
}
// @lc code=end

func main(){
	a := []int{0,1,2,2,3,0,4,2}
	size := removeElement(a, 2)
	fmt.Printf("size is %d, %v", size, a)

	a1 := []int{2}
	size = removeElement(a1, 2)
	fmt.Printf("size is %d, %v", size, a1)


	a3 := []int{2,2,2}
	size = removeElement(a3, 2)
	fmt.Printf("size is %d, %v", size, a3)

}
