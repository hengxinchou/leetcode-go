/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int			
	a = []int{1,2,3,1}
	fmt.Printf("%v, %t\n", a, containsDuplicate(a))

	a = []int{1,2,3,4}
	fmt.Printf("%v, %t\n", a, containsDuplicate(a))

	a = []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Printf("%v, %t\n", a, containsDuplicate(a))
}

func containsDuplicate(nums []int) bool {
	map1 := map[int]bool{}
	for i := 0 ; i < len(nums) ; i++{
		if _, ok := map1[nums[i]]; ok {
			return true
		} else {
			map1[nums[i]] = true
		}
	}
	return false
}
// @lc code=end

