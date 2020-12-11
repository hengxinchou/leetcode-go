/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
package main

import "fmt"

func main() {

	fmt.Printf("%d, %d, %v\n", 4, 5, combine(4, 4))
}

func combine(n int, k int) (ans [][]int) {

	items := []int{}
	for i := 1; i <= n; i++ {
		items = append(items, i)
	}

	set := []int{}
	var dfs func(int)

	dfs = func(j int) {
		if len(set) == k {
			ans = append(ans, append([]int{}, set...))
			return
		} else if j == len(items) {
			return
		}
		set = append(set, items[j])
		dfs(j + 1)
		set = set[:len(set)-1]
		dfs(j + 1)
	}
	dfs(0)
	return
}

// @lc code=end
