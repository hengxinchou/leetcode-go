/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
package main

import "fmt"

func main() {
	a := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}	

	fmt.Printf("%d\n", islandPerimeter(a))
}
func islandPerimeter(grid [][]int) int {
	bottom := len(grid) - 1
	// root := 0
	// left := 0 
	right := len(grid[0]) - 1
	roundSize := 0
	fmt.Printf("bottom is %d, right is %d\n", bottom, right)
	for i, a := range grid {
		for j, _ := range a {
			if grid[i][j] == 0 {
				continue
			}
			count := 0 //
			fmt.Printf("%d, %d, ", i, j) 
			if i == 0 || (i > 0 && grid[i-1][j] == 0 ){
				// if i > 0 {
				// 	fmt.Printf("grid[i-1][j] is %d, ", grid[i-1][j])
				// }
				count++
			}
			if i == bottom || (i < bottom && grid[i+1][j] == 0) {
				// if i < bottom {
				// 	fmt.Printf("grid[i+1][j] is %d, ", grid[i+1][j])
				// }
				
				count++
			}
			
			if j == 0 || (j > 0 && grid[i][j-1] == 0 ){
				// if j > 0 {
				// 	fmt.Printf("grid[i][j-1] is %d, ", grid[i][j-1])
				// }
				count++
			}
			if j == right || ( j < right && grid[i][j+1] == 0) {
				// if j < right {
				// 	fmt.Printf("grid[i][j+1]  is %d, ", grid[i][j+1] )
				// }
				count++
			}

			roundSize += count
			// fmt.Printf("count is %d\n", count)
		}
	}
	return roundSize
}
// @lc code=end

