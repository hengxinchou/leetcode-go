/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a [][]byte
	a = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'}, 
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}		
	for i :=  0 ; i < len(a) ;i ++ {
		for j := 0 ; j < len(a[0]) ; j ++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}

	solve(a)

	fmt.Println("after")
	for i :=  0 ; i < len(a) ;i ++ {
		for j := 0 ; j < len(a[0]) ; j ++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}

	a =[][]byte{{'X','O'},{'O', 'X'}}
	for i :=  0 ; i < len(a) ;i ++ {
		for j := 0 ; j < len(a[0]) ; j ++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}

	solve(a)

	fmt.Println("after")
	for i :=  0 ; i < len(a) ;i ++ {
		for j := 0 ; j < len(a[0]) ; j ++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}
}

func solve(board [][]byte)  {

	if len(board) <= 0 {
		return
	}

	m := len(board)
	n := len(board[0])

	var dfs func (int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}

		board[i][j] = 'A'

		dfs(i - 1, j)
		dfs(i+1, j)
		dfs(i ,j -1)
		dfs(i, j + 1)
	}

	for i := 0 ; i < m ; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 1 ; i < n -1 ; i ++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	for i := 0 ;i < m  ; i++ {
		for j := 0 ; j < n ; j ++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end

