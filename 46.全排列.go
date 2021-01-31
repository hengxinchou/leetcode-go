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
	a = []int{1, 2}
	fmt.Printf("%v, %v\n", a, permute(a))

	a = []int{1, 2, 3}
	fmt.Printf("%v, %v\n", a, permute(a))

	a = []int{5, 4, 2, 6}
	fmt.Printf("%v, %v\n", a, permute(a))
}

func permute1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
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

func permute(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	var dfs func(int, []int)
	dfs = func(pilot int, tmp []int) {
		if pilot == n-1 {
			result = append(result, tmp)
			return
		}
		for i := pilot; i < n; i++ {
			tmp[pilot], tmp[i] = tmp[i], tmp[pilot]
			combination := append([]int{}, tmp...)
			dfs(pilot+1, combination)
			tmp[pilot], tmp[i] = tmp[i], tmp[pilot] // 恢复原来的位置
		}
	}
	dfs(0, nums)
	return result
}

// @lc code=end
