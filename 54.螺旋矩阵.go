/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	rows := len(matrix)
	columns := len(matrix[0])
	target := rows * columns
	a := make([]int, target)
	index := 0
	l, t := 0, 0
	r, b := columns-1, rows-1
	for index < target {
		for i := l; i <= r && index < target ; i++ {
			a[index] = matrix[l][i]
			index++
		}

		t++
		for i := t; i <= b && index < target; i++ {
			a[index] = matrix[i][r]
			index++
		}
		r--
		for i := r; i >= l && index < target ; i-- {
			a[index] = matrix[b][i]
			index++
		}
		b--
		for i := b; i >= t && index < target; i-- {
			a[index] = matrix[i][l]
			index++
		}
		l++
	}
	return a
}

// @lc code=end

func main() {
	// n := 4
	// a := make([][]int, n)
	// for i := 0 ; i < 4 ; i++{
	// 	a[i] = make([]int, 4)
	// }
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// c := make([][]int, 4)
	// a = make([]int, 4)
	// a[0] 
	b := spiralOrder(a)
	fmt.Println(b)

	c := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	d := spiralOrder(c)
	fmt.Println(d)
}
