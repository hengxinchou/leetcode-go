/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
package main

import "fmt"

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := []int{}
	a := map[int]int{}
	for _, i := range nums {
		a[i]++
	}
	limit := len(nums) / 3
	for k, v := range a {
		fmt.Printf("map the count of %d is %v\n", k, v)
		if v > limit {
			res = append(res, k)
		}
	}

	return res
}

func main(){
	a1 := []int{3,2,3}
	fmt.Printf("%v, %v\n", a1, majorityElement(a1))

	a2 := []int{1}
	fmt.Printf("%v, %v\n", a2, majorityElement(a2))

	a3 := []int{1,1,1,3,3,2,2,2}
	fmt.Printf("%v, %v\n", a3, majorityElement(a3))
}
// @lc code=end

