/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int

	a = []int{1, 1, 2}
	fmt.Printf("%v, %v\n", a, permuteUnique(a))

	a = []int{2, 2, 1, 1}
	fmt.Printf("%v, %v\n", a, permuteUnique(a))
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums) // 排序
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(pilot int) {
		if pilot == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}
			perm = append(perm, nums[i])
			vis[i] = true
			backtrack(pilot + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 未完成
func permuteUnique2(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	var dfs func(int, []int)
	dfs = func(pilot int, tmp []int) {
		if pilot == n {
			result = append(result, tmp)
			return
		}
		map1 := map[int]int{}
		for i := pilot; i < n; i++ {
			if _, ok := map1[nums[i]]; ok {
				continue
			}
			tmp[i], tmp[pilot] = tmp[pilot], tmp[i]

			combination := append([]int{}, tmp...)

			dfs(pilot+1, combination)
			tmp[i], tmp[pilot] = tmp[pilot], tmp[i]
			map1[nums[i]] = 1
		}
	}
	dfs(0, nums)
	return result
}

func permuteUnique1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		if nums[0] != nums[1] {
			return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
		} else {
			return [][]int{nums}
		}

	}

	res := [][]int{}
	isDone := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := isDone[nums[i]]; ok {
			continue
		} else {
			isDone[nums[i]] = true
		}

		newA := make([]int, len(nums))
		copy(newA, nums)
		newA[0], newA[i] = newA[i], newA[0]

		res2 := permuteUnique(newA[1:])
		// fmt.Printf("one round, newA[0] is %d, ", newA[0])
		for _, tmp := range res2 {
			tmp2 := append([]int{newA[0]}, tmp...)
			// fmt.Printf("%v, ", tmp2)
			res = append(res, tmp2)
		}
		// fmt.Println()
	}
	return res
}

// @lc code=end
