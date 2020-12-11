/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int 

	a = []int{1,2,3,1}
	fmt.Printf("%v, %t\n", a, containsNearbyDuplicate(a, 3))	
	a = []int{1,0,1,1}
	fmt.Printf("%v, %t\n", a, containsNearbyDuplicate(a, 1))	
	a = []int{1,2,3,1,2,3}
	fmt.Printf("%v, %t\n", a, containsNearbyDuplicate(a, 2))	
}

func containsNearbyDuplicate(nums []int, k int) bool {
	map1 := map[int]int{}
	for i := 0 ; i < len(nums) ; i++ {
		if v, ok := map1[nums[i]]; ok  {
			if i - v <= k {
				return true
			}
		} 
	    map1[nums[i]] = i
	}
	return false
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
// @lc code=end

