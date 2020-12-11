/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
package main

import "fmt"

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num % 2 == 0 {
		num = num / 2
	}
	for num % 3 == 0 {
		num = num / 3
	}
	for num % 5 == 0 {
		num = num / 5
	}
	return num == 1
}

// @lc code=end

func main(){
	fmt.Printf("%d, %t\n", 6, isUgly(6))
	fmt.Printf("%d, %t\n", 8, isUgly(8))
	fmt.Printf("%d, %t\n", 14, isUgly(14))
	fmt.Printf("%d, %t\n", 1, isUgly(1))
}