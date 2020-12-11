/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
package main

import "fmt"

func main() {

	var nums []int

	nums = []int{1,2,3,4,5}
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	nextPermutation(nums)
	fmt.Printf("%v, %v\n", nums2, nums)
}

func nextPermutation(nums []int)  {
	nums[0] = 0
}
// @lc code=end

