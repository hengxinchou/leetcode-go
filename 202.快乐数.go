/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start

// @lc code=end

package main

import "fmt"
// import "strconv"

func isHappy(n int) bool {
	// sum := 0
	if n < 1 {
		return false
	}
	map1 := map[int]bool{}
	for n != 1 {
		a := []int{}
		for n != 0 {
			tmp := n % 10
			a = append([]int{tmp}, a...)
			n = n / 10 
		}
		sum2 := 0 
		fmt.Println(a)
		for _, v := range a {
			sum2 += sum2 * 10 + v
		}
		if _, ok := map1[sum2]; ok {
			return false
		} else{
			map1[sum2] = true
		}
		sum := 0
		 
		for _, b := range a {
			sum += b * b
		}
		n = sum
	}
	return true
}
func main() {
	a := 21112
	fmt.Printf("%d, %t\n", a, isHappy(a))	
}
