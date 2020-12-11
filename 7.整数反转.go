/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
package main

import "math"
import "fmt"
// import "strconv"

func reverse(x int) int {
	negative := false
	if x < 0  {
		negative = true
		x = -x
	}
	res := 0
	left := x / 10 
	more := x % 10
	for more != 0 || left > 0 {
		res = res * 10 + more
		x = left
		left = x / 10 
		more = x % 10

		if res > math.MaxInt32 {
			return 0
		}
	} 
	if negative {
		res = -res
	}
	return res
}
// @lc code=end

func main(){
	i1 := 123
	fmt.Printf("%d, %d\n", i1, reverse(i1))

	i1 = -123
	fmt.Printf("%d, %d\n", i1, reverse(i1))


	i1 = 120
	fmt.Printf("%d, %d\n", i1, reverse(i1))

	i1 = 0
	fmt.Printf("%d, %d\n", i1, reverse(i1))

	i1 = 1534236469
	fmt.Printf("%d, %d\n", i1, reverse(i1))
}

