/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int
	a = []int{1,2,2,1,1,3}
	fmt.Printf("%v, %t\n", a, uniqueOccurrences(a))	

	a = []int{1,2}
	fmt.Printf("%v, %t\n", a, uniqueOccurrences(a))	

	a = []int{-3,0,1,-3,1,1,1,-3,10,0}
	fmt.Printf("%v, %t\n", a, uniqueOccurrences(a))	
}

func uniqueOccurrences(arr []int) bool {
	tmp := map[int]int{}
	for _, i := range arr {
		tmp[i]++
	}
	tmp2 := map[int]int{}
	for _, count := range tmp {
		tmp2[count]++
		if tmp2[count] > 1 {
			return false
		}
	}
	return true
}
// @lc code=end

