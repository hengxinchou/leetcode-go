/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int
	a = []int{1,2,3}
	
	fmt.Printf("%v\n", permute(a))

	a = []int{5,4,2,6}
	fmt.Printf("%v\n", permute(a))
}

func permute(nums []int) [][]int{
	if len(nums) == 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	res := [][]int{}
	for i := 0 ; i < len(nums) ; i++ {
		newA := make([]int, len(nums))
		copy(newA, nums)
		newA[0], newA[i] = newA[i], newA[0]
		
		res2 := permute(newA[1:])
		for _, tmp := range res2 {
			tmp2 := append([]int{newA[0]}, tmp...)
			res = append(res, tmp2)
		}
	}
	return res
}
// @lc code=end

