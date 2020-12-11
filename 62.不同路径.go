/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
package main

import "fmt"

func main() {
	
	fmt.Println(uniquePaths(3,2))
	fmt.Println(uniquePaths(7,3))
	fmt.Println(uniquePaths(23,12))
}

// type Point struct {
	// i,j int
// }

func uniquePaths2(m int, n int) int {
	paths := 0 	
	// map1 = map[Point]int
	var dfs func (int, int) 

	dfs = func(i, j int) {
		if i == m - 1 && j == n - 1 {
			paths += 1
			return
		}
		if i == m {
			return
		} 
		if j == n {
			return
		}

		dfs(i, j+1)
		dfs(i+1, j)
	}
	dfs(0, 0)
	return paths
}
// @lc code=end


func uniquePaths(m, n int) int {

	states := make([][]int, n)
	for i := 0 ; i < n ; i++ {
		states[i] = make([]int, m)
	}

	for j := 0 ; j < m ; j++ {
		states[0][j] = 1
	}
	for i := 0 ; i < n ; i ++ {
		states[i][0] = 1
	}

	for i := 1 ; i < n ; i++ {
		for j := 1; j < m ; j++ {
			states[i][j] = states[i-1][j] + states[i][j-1]
		}
	}
	return states[n-1][m-1]
}
