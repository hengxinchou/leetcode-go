/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 
	}

	down := 0
	upper := len(nums) -1
	var mid int

	for  down <= upper {
		mid = down+((upper-down)>>1)
		// fmt.Printf("%d, %d, %d\n", down, upper, mid)
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			down = mid + 1 
		} else {
			upper = mid - 1
		}
	}
	return -1
}
// @lc code=end






func main(){
	nums := []int{-1,0,3,5,9,12}
	target := -1
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))
	target = 12
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))
	target = 5
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))

	target = -1
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))

	target = 13
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))
	target = -2
	fmt.Printf("%v, %d, %d\n",  nums, target, search(nums, target))

}
