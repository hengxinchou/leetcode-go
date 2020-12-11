/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	i := 0 
	first := strs[0]

	for {
		if i >=  len(first) {
			return first[0:i]
		}
		s := first[i]
		for j := 1 ; j < len(strs) ; j++ {
			// fmt.Printf("i is %d, s is %s, strs[%d] is %s\n", i, string(s), j, strs[j])
			if i >= len(strs[j]) || s != strs[j][i] {
				return first[0:i]
			}
		}
		i++
	}
	return first
}
// @lc code=end

func main(){
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))

	fmt.Println(longestCommonPrefix([]string{"fl","flower","flight"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix(nil))
}
