/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
package main

import "fmt"

func main() {

	var a []int
	a = []int{1, 2, 3, 1}
	fmt.Printf("%v, %d\n", a, findPeakElement(a))

	a = []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Printf("%v, %d\n", a, findPeakElement(a))
}

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	n := len(nums)
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

// @lc code=end
