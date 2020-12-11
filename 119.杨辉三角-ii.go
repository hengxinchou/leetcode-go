/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
package main 

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := [][]int{[]int{1}}

	for i := 1 ; i <= rowIndex ; i ++ {
		prev := res[i-1]
		tmp := []int{}
		tmp = append(tmp, 1)
		for j := 1 ; j < i; j++ {
			tmp = append(tmp, prev[j-1]+prev[j])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res[rowIndex]
}
// @lc code=end

func main(){
	fmt.Println(getRow(3))
}