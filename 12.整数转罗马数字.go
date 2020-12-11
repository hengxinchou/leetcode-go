/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
package main

import "fmt"
import "math"
import "strconv"

func main() {
	var a int 
	a = 90
	fmt.Printf("%d, %s\n", a, intToRoman(a))

	a = 1049
	fmt.Printf("%d, %s\n", a, intToRoman(a))

	a = 100049
	fmt.Printf("%d, %s\n", a, intToRoman(a))
}

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	size := len(intSlice) - 1
	res := ""
	for num != 0 {
		a := strconv.Itoa(num)
		x := len(a) - 1
		b :=  (num / int(math.Pow10(x))) * int(math.Pow10(x))
		for size >= 0 {
			if b >= intSlice[size] {
				res += roman[size]	
				num -= intSlice[size]
				break
			} else {
				size--
			}
		}
	}
	return res
}
// @lc code=end

