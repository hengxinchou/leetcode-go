/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
package main

import "fmt"

func main() {
	a := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}	
	fmt.Println(minPathSum(a))
}

func minPathSum(grid [][]int) int {
  n := len(grid)
  m := len(grid[0])

  states := make([][]int, n)
  for i := 0 ; i < n ; i++ {
	  states[i] = make([]int, m)
  }

  sum := 0
  for j := 0 ; j < m ; j++ {
	  sum += grid[0][j]
	  states[0][j] = sum
  }

  sum = 0
  for i := 0 ; i < n ; i++ {
	  sum += grid[i][0]
	  states[i][0] = sum
  }

  for i := 1; i < n ; i++ {
	  for j := 1 ; j < m ; j++ {
		  states[i][j] = min(states[i-1][j], states[i][j-1]) + grid[i][j]
	  }
  }
  return states[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
// @lc code=end

