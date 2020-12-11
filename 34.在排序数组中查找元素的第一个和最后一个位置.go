/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int
	a = []int{5,7,7,8,8,10}	
	fmt.Printf("%v, %v\n", a, searchRange(a, 8))

	a = []int{5,7,7,8,8,10}	
	fmt.Printf("%v, %v\n", a, searchRange(a, 6))
}

func searchRange(nums []int, target int) []int {
	a := search(nums, target)
	if a == -1 {
		return []int{-1,-1}
	}

	// fmt.Printf("a is %d\n", a)
	start := a
	for start >= 0 && nums[start] == target  {
		start--
	}
	// fmt.Printf("start is %d\n", start)
	if start < 0 {
		start = 0
	} else {
		start++
	}
	
	end := a
	for end < len(nums) && nums[end] == target {
		end++
	}
	// fmt.Printf("end is %d\n", end)
	if end == len(nums) {
		end = len(nums) - 1
	} else {
		end--
	}
	return []int{start, end}
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 
	}

	down := 0
	upper := len(nums) -1
	var mid int

	for  down <= upper {
		mid = down+((upper-down)>>1)
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

