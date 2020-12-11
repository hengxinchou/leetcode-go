/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	a = "25525511135"
	fmt.Printf("%s, %v\n", a, restoreIpAddresses(a))

	a = "0000"
	fmt.Printf("%s, %v\n", a, restoreIpAddresses(a))

	a = "1111"
	fmt.Printf("%s, %v\n", a, restoreIpAddresses(a))

	a = "101023"
	fmt.Printf("%s, %v\n", a, restoreIpAddresses(a))

	a = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	fmt.Printf("%s, %v\n", a, restoreIpAddresses(a))
}

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}

	if len(s) > 12 {
		return []string{}
	}

	// size := len(s)
	res := []string{}

	// for i := 0 ; i < len(s) - 3 ; i++ {
	for j := 1; j < len(s)-2; j++ {
		for k := j + 1; k < len(s)-1; k++ {
			for l := k + 1; l < len(s); l++ {
				tmp := []string{s[0:j], s[j:k], s[k:l], s[l:]}
				// fmt.Printf("tmp is %v, %d, %d, %d, %d\n", tmp, i, j, k, l)
				if validArray(tmp) {
					res = append(res, strings.Join(tmp, "."))
				}
			}
		}
	}
	// }
	return res
}

func validArray(a []string) bool {
	// fmt.Printf("%v, size is %d\n", a, len(a))
	for i := 0; i < len(a); i++ {
		if !valid(a[i]) {
			return false
		}
	}
	return true
}

func valid(s string) bool {
	// fmt.Printf("valid s is %s\n", s)
	if s[0] == '0' && len(s) > 1 {
		return false
	}

	num, _ := strconv.Atoi(s)
	if num >= 0 && num <= 255 {
		return true
	} else {
		return false
	}
}

// @lc code=end
