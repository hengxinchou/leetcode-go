/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
package main

import "fmt"

func main() {
	var n int
	n = 4
	fmt.Printf("%d, %d\n", n, totalNQueens(n))
	n = 8
	fmt.Printf("%d, %d\n", n, totalNQueens(n))
	n = 7
	fmt.Printf("%d, %d\n", n, totalNQueens(n))
	n = 6
	fmt.Printf("%d, %d\n", n, totalNQueens(n))
}

func totalNQueens(n int) int {
	count := 0
	result := make([]int, n)	
	var dfs func (int, int) 

	dfs = func(row int, max int) {
		if row == max {
			count++
			return
		}
		for column := 0 ; column < max ; column++ {
			if isOk(row, column, max, result) {
				result[row] = column
				dfs(row+1, max)
			}
		}
	}
	dfs(0, n)
	return count
}

func isOk(row, column, max int, result []int) bool {
	leftUp := column - 1
	rightUp := column + 1	
	for i := row - 1 ; i >= 0 ; i-- {
		if result[i] == column {
			return false
		}
		if leftUp >= 0 {
			if result[i] == leftUp  {
				return false
			}
		}
		if rightUp < max {
			if result[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}

func printString(result []int, max int) []string {
	res := []string{}
	for i := 0 ; i < max ; i ++ {
		tmp := ""
		for j := 0 ; j < max ; j++ {
			if result[i] == j {
				tmp = tmp + "Q"
			} else {
				tmp = tmp + "."
			}
		}
		res = append(res, tmp)
	}
	return res
}
// @lc code=end

