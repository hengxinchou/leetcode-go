/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
package main

import "fmt"

func main() {

	var text, pattern string

	text = "cb"
	pattern = "?a"
	fmt.Printf("%s, %s, %t\n",text, pattern, isMatch(text, pattern))

	text = "adcb"
	pattern = "*a*b"
	fmt.Printf("%s, %s, %t\n",text, pattern, isMatch(text, pattern))

	text = "acdcb"
	pattern = "a*c?b"
	fmt.Printf("%s, %s, %t\n",text, pattern, isMatch(text, pattern))

	text = "aa"
	pattern = "*"
	fmt.Printf("%s, %s, %t\n",text, pattern, isMatch(text, pattern))

	text ="aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba"
	pattern = "a*******b"
	fmt.Printf("%s, %s, %t\n",text, pattern, isMatch(text, pattern))
}

func isMatch(text string, pattern string) bool {
	matched := false
	var dfs func (int,int) 
	dfs = func(ti, pj int) {
		if matched {
			return
		}
		if pj == len(pattern) {
			if ti == len(text) { // 文本串也到结尾了
				matched = true
			}
			return
		}

		if pattern[pj] == '*' { // *匹配任意个字符
			for k := 0 ; k <= len(text) - ti ;  k++ {
				dfs(ti + k, pj + 1)
			}
		} else if pattern[pj] == '?' { // ?匹配0个或者1个字符
			// dfs(ti, pj + 1)
			dfs(ti + 1, pj + 1)
		} else if ti < len(text) && pattern[pj] == text[ti] { // 纯字符匹配才行
			dfs(ti+1, pj+1)
		}

	}
	dfs(0, 0)
	return matched	
}
// @lc code=end

