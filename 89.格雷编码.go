/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
package main

import "fmt"

func main() {
	fmt.Printf("%v\n", grayCode(0))
	fmt.Printf("%v\n", grayCode(1))
	fmt.Printf("%v\n", grayCode(2))
	fmt.Printf("%v\n", grayCode(3))
	fmt.Printf("%v\n", grayCode(4))
}

func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head + res[j]);
		}	
		head <<= 1
	}
	return res

	// if n < 0 {
	// 	return []int{}
	// }
	// if n == 0 {
	// 	return []int {0}
	// }
	// a := grayCode2(n)
	// fmt.Printf("%v\n", a)
	// res := []int{}
	// for i := 0 ; i < len(a) ; i++ {
	// 	res = append(res, transfer(a[i]))
	// }
	// return res
}

func transfer(num string) int {
	sum := 0
	for i := 0 ; i < len(num) ; i++ {
		sum = sum *2 +  int(num[i] - '0')
	}
	return sum
}
func grayCode2(n int) []string{
	if n == 1  {
		return []string{"0", "1"}
	}
	a1 := grayCode2(n-1)
	// a2 := grayCode2(n-1, false)
	res := []string{}
	for i := 0 ; i < len(a1) ; i++ {
		res = append(res, "0" + a1[i])
	}
	for i := len(a1) -1 ; i >= 0 ; i-- {
		res = append(res, "1" + a1[i])
	}
	return res
}
// @lc code=end

