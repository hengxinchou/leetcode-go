/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
package main

import "fmt"

func main() {
	var s, p string

	s = "aa"
	p = "a"
	fmt.Printf("%s, %s, %t\n", s, p, isMatch(s, p))

	s = "aa"
	p = "a*"
	fmt.Printf("%s, %s, %t\n", s, p, isMatch(s, p))

	s = "ab"
	p = ".*"
	fmt.Printf("%s, %s, %t\n", s, p, isMatch(s, p))

	s = "aab"
	p = "c*a*b"
	fmt.Printf("%s, %s, %t\n", s, p, isMatch(s, p))

	s = "mississippi"
	p = "mis*is*p*."
	fmt.Printf("%s, %s, %t\n", s, p, isMatch(s, p))
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		// . 必然匹配， 任意一个字符
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	// m是字符串的长度
	states := make([][]bool, m+1)
	// n 是 正则的长度
	for i := 0; i < len(states); i++ {
		states[i] = make([]bool, n+1) // 默认为false
	}
	states[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ { // 从1开始，而不是从0 开始
			if p[j-1] == '*' { // 如果pattern 前一个是 *
				states[i][j] = states[i][j] || states[i][j-2]
				if matches(i, j-1) {
					states[i][j] = states[i][j] || states[i-1][j]
				}
			} else if matches(i, j) { // 如果匹配再进行下面的赋值
				states[i][j] = states[i][j] || states[i-1][j-1]
			}
		}
	}
	return states[m][n]
}

func isMatch1(s string, p string) bool {
	if s == p {
		return true
	}

	if len(p) > 1 && p[1] == '*' {
		switch p[0] {
		case '.':
			for i := 0; i <= len(s); i++ {
				if isMatch(s[i:], p[2:]) {
					return true
				}
			}
		default:
			if isMatch(s, p[2:]) {
				return true
			}

			for i := 0; i < len(s) && s[i] == p[0]; i++ {
				if isMatch(s[i+1:], p[2:]) {
					return true
				}
			}
		}

	} else {
		if s == "" || p == "" {
			return false
		}

		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}
	}
	return false
}

// @lc code=end
