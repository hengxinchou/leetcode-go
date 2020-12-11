/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
package main
import "fmt"
func main() {
	a := [][]int {
		{0,0,1},
		{0,0,0},
		{0,0,0},
	}
	fmt.Println(uniquePathsWithObstacles(a))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n :=  len(obstacleGrid)
	m :=  len(obstacleGrid[0])

	states := make([][]int, n)

		for i := 0 ; i < n ; i++ {
			states[i] = make([]int, m)
		}

		for j := 0 ; j < m ; j++ {
			if obstacleGrid[0][j] == 1 {
			    states[0][j] = 0
			} else if j == 0 {
				states[0][j] = 1
			} else  {
				states[0][j] = states[0][j-1]
			} 
		}
		for i := 0 ; i < n ; i ++ {
			if obstacleGrid[i][0] == 1 {
				states[i][0] = 0
			} else if i == 0 {
				states[i][0] = 1
			} else {
				states[i][0] = states[i-1][0]
			}
		}

		for i := 1 ; i < n ; i++ {
			for j := 1; j < m ; j++ {
				if obstacleGrid[i][j] == 1 {
					states[i][j] = 0
				} else {
					states[i][j]=states[i-1][j] + states[i][j-1]
				}
			}
		}
		return states[n-1][m-1]
}
// @lc code=end

