/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
package main

import "fmt"

func majorityElement(nums []int) int {
	a := map[int]int{}
	for _, i := range nums {
		a[i]++
	}
	limit := len(nums) / 2
	fmt.Printf("limit is %d\n", limit)
	for k, v := range a {
		fmt.Printf("map the count of %d is %v\n", k, v)
		if v > limit {
			return k
		}
	}

	return  -1
}
// @lc code=%end
func main(){
	a1 := []int{3,2,3}
	fmt.Printf("%v, %v\n", a1, majorityElement(a1))

	a2 := []int{1}
	fmt.Printf("%v, %v\n", a2, majorityElement(a2))

	a3 := []int{2,2,1,1,1,2,2}
	fmt.Printf("%v, %v\n", a3, majorityElement(a3))
}
