/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func main() {
	var triangle [][]int

	triangle = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println("result")
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n <= 0 {
		return 0
	}

	states := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Println(triangle[i])
		states[i] = make([]int, len(triangle[i]))
	}

	for i := 0; i < n; i++ {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			if i == 0 {
				states[i][j] = triangle[i][j]
				continue
			}
			if j == 0 {
				states[i][j] = states[i-1][0] + triangle[i][j]
				continue
			}

			if j == m-1 {
				states[i][j] = states[i-1][len(triangle[i-1])-1] + triangle[i][j]
				continue
			}
			states[i][j] = min(states[i-1][j-1], states[i-1][j]) + triangle[i][j]
		}
	}

	minV := math.MaxInt32

	final := n - 1
	for i := 0; i < len(states[final]); i++ {
		if states[final][i] < minV {
			minV = states[final][i]
		}
	}
	return minV
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// @lc code=end
