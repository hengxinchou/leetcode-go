/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
package main

import "fmt"
import "math"

// func multiply(num1 string, num2 string) string {
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	sum1, sum2 := 0, 0
// 	for i := 0 ; i < len(num1) ; i++{
// 		sum1 = sum1 * 10 + int(num1[i] - '0') 
// 	}
// 	for i := 0 ; i < len(num2) ; i++{
// 		sum2 = sum2 * 10 + int(num2[i] - '0')
// 	}

// 	sum := int64(sum1) * int64(sum2)
// 	fmt.Printf("%d, %d, muiltply is %d\n", sum1, sum2, sum)
// 	res := ""
// 	for sum != 0 {
// 		tmp := sum % 10
// 		res = string('0'+ tmp) + res
// 		sum = sum / 10
// 	}
// 	return res
// }

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    m, n := len(num1), len(num2)
    ansArr := make([]int, m + n)
    for i := m - 1; i >= 0; i-- {
        x := int(num1[i]) - '0'
        for j := n - 1; j >= 0; j-- {
            y := int(num2[j] - '0')
            ansArr[i + j + 1] += x * y
        }
    }
    for i := m + n - 1; i > 0; i-- {
        ansArr[i - 1] += ansArr[i] / 10
        ansArr[i] %= 10
    }
    ans := ""
    idx := 0
    if ansArr[0] == 0 {
        idx = 1
    }
    for ; idx < m + n; idx++ {
        ans += strconv.Itoa(ansArr[idx])
    }
    return ans
}
// @lc code=end

func main(){
	fmt.Printf("max32 is %d, max64 is %d\n", math.MaxInt32, math.MaxInt64)
	a , b := "2", "3"
	fmt.Printf("%s, %s, %s\n", a, b, multiply(a, b))

	a , b = "0", "3"
	fmt.Printf("%s, %s, %s\n", a, b, multiply(a, b))

	a , b = "123", "456"
	fmt.Printf("%s, %s, %s\n", a, b, multiply(a, b))

	a, b = "498828660196", "840477629533"
	fmt.Printf("%s, %s, %s\n", a, b, multiply(a, b))
}

