/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func main() {
	var dividend, divisor int
	dividend = 10
	divisor = 3
	fmt.Printf("%d, %d, %d\n", dividend, divisor, divide(dividend, divisor))

	dividend = 7
	divisor = -3
	fmt.Printf("%d, %d, %d\n", dividend, divisor, divide(dividend, divisor))

	dividend = -2147483648
	divisor = 1
	fmt.Printf("%d, %d, %d\n", dividend, divisor, divide(dividend, divisor))

	dividend = 10
	divisor = 3
	fmt.Printf("%d, %d, %d\n", dividend, divisor, divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {

	flag := false

	if dividend < 0 {
		flag = !flag
		dividend = (-1) * dividend
	}

	if divisor < 0 {
		divisor = (-1) * divisor
		flag = !flag
	}

	count := div(dividend, divisor)

	if flag {
		count = (-1) * count
	}

	if count > math.MaxInt32 {
		return math.MaxInt32
	} else if count < math.MinInt32 {
		return math.MinInt32
	} else {
		return count
	}
}

// all is positive
func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	c := b
	for (c + c) < a {
		count = count + count
		c = c + c
	}
	return count + div(a-c, b)
}

// @lc code=end
