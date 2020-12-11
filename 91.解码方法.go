/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
package main
import "fmt"
func main() {
	
	var s string
	s = "226"
	fmt.Println(numDecodings(s))

	s = "12"
	fmt.Println(numDecodings(s))
}

func numDecodings2(s string) int {
	// res := 0
	// var dfs(int, int)
	// dfs = func(i, j int) {
	// 	if i == len(s)  {
			
	// 		return
	// 	}
	// 	dfs(0, 1)	
	// }
	// dfs(0,1)
	return 0
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	p, q := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' && s[i-1] != '1' && s[i-1] != '2' {
			return 0
		}
		if s[i] == '0' {
			q = p
			continue
		}
		if (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1') {
			p, q = q, p + q	
			continue
		} 
		p = q
	}
	return q
}
// @lc code=end

