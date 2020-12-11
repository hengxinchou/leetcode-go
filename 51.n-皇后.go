/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
package main

import "fmt"

func main() {
	n := 4
	allResult := solveNQueens(n)
	for i := 0 ; i < len(allResult) ; i++ {
		for j := 0 ; j < n ; j++ {
			fmt.Println(allResult[i][j])
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	allResult := [][]string{}
	result := make([]int, n)	
	var dfs func (int, int) 

	dfs = func(row int, max int) {
		if row == max {
			tmp := printString(result, max)
			allResult = append(allResult, tmp)
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
	return allResult
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

