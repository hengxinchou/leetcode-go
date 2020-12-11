/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	i := n - 1
	digits[i] = digits[i] + 1
	for digits[i] >= 10 && i > 0 {
		digits[i] = digits[i] - 10
		i-- 
		digits[i]++
	}
	if i == 0 && digits[i] >= 10 {
		digits[i] = digits[i] - 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
// @lc code=end

func main(){
	a1 := []int{1,2,3}
	fmt.Printf("%d, %d\n", a1, plusOne(a1))

	a2 := []int{4,3,2,1}
	fmt.Printf("%d, %d\n", a2, plusOne(a2))

	a3 := []int{9}
	fmt.Printf("%d, %d\n", a3, plusOne(a3))

	a4 := []int{0}
	fmt.Printf("%d, %d\n", a4, plusOne(a4))

	a5 := []int{9,9}
	fmt.Printf("%d, %d\n", a5, plusOne(a5))


}

