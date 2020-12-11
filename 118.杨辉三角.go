/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	res := [][]int{[]int{1}}

	// if numRows == 1 {
	// 	return res
	// }

	for i := 1 ; i < numRows ; i ++ {
		prev := res[i-1]
		tmp := []int{}
		tmp = append(tmp, 1)
		for j := 1 ; j < i  ; j++ {
			tmp = append(tmp, prev[j-1]+prev[j])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res
}
// @lc code=end
func print(a [][]int ){
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println()
}

func main(){
	print(generate(0))
	print(generate(1))
	print(generate(2))
	print(generate(5))
}
