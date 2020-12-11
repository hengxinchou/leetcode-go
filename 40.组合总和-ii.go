/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
package main

import "fmt"
import "sort"

func main() {
	var candidates []int
	var target int

	candidates = []int{10,1,2,7,6,1,5}
	target = 8
	fmt.Printf("%v, target is %d, %v\n", candidates, target, combinationSum(candidates, target))

	candidates = []int{2,5,2,1,2}
	target = 5
	fmt.Printf("%v, target is %d, %v\n", candidates, target, combinationSum(candidates, target))
}

func combinationSum(candidates[]int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	res := [][]int{}
	for index, n := range candidates {
		if index > 0 && n == candidates[index-1] {
			continue
		}
		if n == target {
			res = append(res, []int{n})
			continue
		} else if n > target {
			break
		} else {
			for _, j := range dfs(candidates[index + 1: ], target - n) {
				k := append([]int{n}, j...) // j is array
				res = append(res, k) 
			}
		}
	}
	return res
}
// @lc code=end

