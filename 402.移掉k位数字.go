/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */
package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main() {
	var a string
	var b int
	a = "1432219"
	b = 3
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))

	a = "10200"
	b = 1
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))

	a = "112"
	b = 1
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))

	a = "10"
	b = 1
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))

	a = "1111111"
	b = 3
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))

	a = "1234567890"
	b = 9
	fmt.Printf("%s, %d, the result %s\n", a, b, removeKdigits(a, b))
}
