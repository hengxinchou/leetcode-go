/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
package main

import "fmt"
import "math"

func maxSlidingWindow(nums []int, k int) []int {

	// stack := []int{}
	if len(nums) <= 0 {
		return []int{}
	}

	max, maxPos := math.MinInt32, -1

	if k >= len(nums) {
		max, _ = findMax(nums)
		return []int{max}
	} 

	res := []int{}
	
	// res = append(res, max)
	// low, upper := 0, k
	for i := 0 ; i < len(nums) - k +1 ; i++ {
		low := i
		upper := i + k -1
		fmt.Printf("%d, %d\n", low, upper)
		if  maxPos >= low {
			fmt.Printf("maxPos <= low\n")
			if max < nums[upper] {
				max = nums[upper]
				maxPos = upper
			} 
			res = append(res, max)
		} else {
			max, maxPos = findMax(nums[low: upper+1])
			// fmt.Printf("findMax %d, %d\n", max, maxPos)
			res = append(res, max)
		}
		// fmt.Printf("res is %v\n", res)
	}
	return res
}
// @lc code=end

func findMax(nums []int) ( max int, pos int) {
	max = math.MinInt64
	for k, i := range nums {
		if i > max {
			max = i
			pos = k
		}
	}
	return 
}

func main(){

	a := []int{1,3,-1,-3,5,3,6,7}
	fmt.Printf("%v, %v\n",a, maxSlidingWindow(a, 5))

	a = []int{1}
	fmt.Printf("%v, %v\n",a, maxSlidingWindow(a, 1))
}

