/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
package main

import "fmt"

func main() {
	var s string
	s = "aab"
	fmt.Println("%s, %v\n", s, partition(s))

	s = "aabba"
	fmt.Println("%s, %v\n", s, partition(s))
	// var a string
	// a = "aa"
	// fmt.Printf("%s %t\n", a, Palindrome(a))
	// a = "aab"
	// fmt.Printf("%s %t\n", a, Palindrome(a))
	// a = "aabb"
	// fmt.Printf("%s %t\n", a, Palindrome(a))
	// a = "a"
	// fmt.Printf("%s %t\n", a, Palindrome(a))
	// a = ""
	// fmt.Printf("%s %t\n", a, Palindrome(a))
}

func partition(s string) [][]string {
	final := [][]string{}
	if len(s) <= 1 {
		final = append(final, []string{s})
		return final
	}

	var dfs func(int, []string)
	dfs = func(start int, res []string) {
		if start == len(s) {
			tmp := make([]string, len(res))
			copy(tmp, res)
			final = append(final, tmp)
			return
		}
		for i := start + 1; i <= len(s); i++ {
			tmp := s[start:i]
			if Palindrome(tmp) {
				dfs(i, append(res, tmp))
			}
		}
	}
	dfs(0, []string{})
	return final
}

func Palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	n := len(s)
	mid := n / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
