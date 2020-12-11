/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package main

import "fmt"
// import "sort"

func main() {
	var candidates []int
	var target int

	candidates = []int{2,3,6,7}
	target = 7
	fmt.Printf("%v, %v\n", candidates, combinationSum(candidates, target))

	candidates = []int{2,3,5}
	target = 8
	fmt.Printf("%v, %v\n", candidates, combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if target <= 0 {
		return res
	}
	for index, i := range candidates {
		if i == target {
			res = append(res, []int{i})
		} else if i > target {
			continue
		} else {
			tmp := combinationSum(candidates[index:], target - i)
			// fmt.Printf("tmp is %v, target is %d, i is %d\n", tmp, target -i, i )
			if len(tmp) == 0 {
				continue
			}
			for _, j := range tmp {
				j = append(j, i) // j is array
				res = append(res, j) 
			}
		}
	}
	fmt.Printf("raw is %v, res is %v, target is %d\n",candidates,  res, target)
	return res
}
// @lc code=end



// func combinationSum(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	return dfs(candidates, target)
// }

// func dfs(candidates []int, target int) [][]int {
// 	ret := [][]int{}
// 	for i, d := range candidates {
// 		if target-d < 0 {
// 			break
// 		} else if target-d == 0 {
// 			ret = append(ret, []int{d})
// 			continue
// 		}
// 		for _, v := range dfs(candidates[i:], target-d) {
// 			ret = append(ret, append([]int{d}, v...))
// 		}
// 	}
// 	return ret
// }