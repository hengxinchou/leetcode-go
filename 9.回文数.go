/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	reverse := 0
	for y > 0 {
		reverse = reverse*10 + y%10
		y = y / 10
	}
	return reverse == x
}

// @lc code=end

func main() {
	fmt.Printf("-121 is  huiwen? %t\n", isPalindrome(-121))
	fmt.Printf("121 is  huiwen? %t\n", isPalindrome(121))
	fmt.Printf("123 is  huiwen? %t\n", isPalindrome(123))
	fmt.Printf("1234321 is  huiwen? %t\n", isPalindrome(1234321))
	// a := []int{}
	// for i := 0 ; i < 10 ; i++{
	// 	a[i] = i
	// }
	// n := 10
	// var a [n]int
	// a := 10000
	// fmt.Println(len(a))
}
