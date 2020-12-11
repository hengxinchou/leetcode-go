/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start

// @lc code=end

package main

import "fmt"

func trailingZeroes(n int) int {

	// five := 0 
	// zero := 0
	// fives := 
	// return n / 5
	sum := 0 
	for n > 0 {
		sum += n / 5
		n = n / 5
	}
	return sum
}

func sum(n int) int {
	sum := 1 
	return sum
	// for i := 1 ; i <= n ; i++{
	// 	sum *=i
	// }
	// return sum
}

func main() {
	var n int

	n = 5
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))

	
	n = 10
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))

	n = 15
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))

	n = 20
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))

	n = 25
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))

	n = 30
	fmt.Printf("%d, %d, %d\n", n, sum(n), trailingZeroes(n))
}