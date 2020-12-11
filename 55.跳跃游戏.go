/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func main() {
	var a []int

	a = []int{2, 3, 1, 1, 4}

	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{3, 2, 1, 0, 4}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{3, 2, 1, 0, 4}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{1, 2}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{1, 0, 1, 0}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{0, 1}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{1, 0, 1, 0}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{0}
	fmt.Printf("%v, %t\n", a, canJump(a))

	a = []int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6}
	fmt.Printf("%v, %t\n", a, canJump(a))
}

func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canJump3(nums []int) bool {

	prev := len(nums) - 2
	end := len(nums) - 1

	// fmt.Printf("nums is %v, prev is %d, end is %d\n", nums, prev, end)
	for prev >= 0 {
		if nums[prev]+prev >= end {
			end = prev
			prev--
		} else {
			prev--
		}
	}
	if nums[0] < end {
		return false
	} else {
		return true
	}
}

func canJump2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	// if nums[0] == 0 {
	// 	return false
	// }
	j := 0
	for ; j < len(nums)-1; j++ {
		if nums[j] == 0 && j < len(nums)-1 {
			return false
		}
		for i := j; i < len(nums)-1; {
			// fmt.Printf("i is %d\n",i)
			if nums[i] == 0 && i < len(nums)-1 {
				break
			}
			// fmt.Printf("i is %d,  i + nums[i] is %d\n", i, i + nums[i])
			if nums[i]+i >= len(nums)-1 {
				return true
			}
			tmp := max2(nums, i, i+nums[i])
			// fmt.Printf("tmp is %d\n", tmp)
			i = tmp + nums[tmp]
		}
	}
	//
	// fmt.Printf("last, i is %d\n", i)
	if j <= len(nums)-1 {
		return false
	} else {
		return true
	}
}

func max2(nums []int, start, end int) int {
	max := math.MinInt32
	pos := start
	for i := start; i < len(nums) && end < len(nums) && i < end; i++ {
		if nums[i]+i >= max {
			max = nums[i] + i
			pos = i
		}
	}
	return pos
}

// @lc code=end
