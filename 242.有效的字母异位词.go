/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
package main
import (
	"fmt"
	// "sort"
	// "strings"
)
func main() {
	var s, t string
	s = "anagram"
	t = "nagaram"
	fmt.Printf("%s, %s, %t\n", s, t,isAnagram(s,t))	

	s = "rat"
	t = "car"
	fmt.Printf("%s, %s, %t\n", s, t,isAnagram(s,t))	

	s = "a"
	t = "ab"
	fmt.Printf("%s, %s, %t\n", s, t,isAnagram(s,t))	
}
func isAnagram(s string, t string) bool {
	// a := strings.Split(s, "")
	// b := strings.Split(t, "")
	// sort.Strings(a)
	// sort.Strings(b)
	// aa := strings.Join(a, "")
	// bb := strings.Join(b, "")
	// return aa == bb
	map1 := map[byte]int{}
	map2 := map[byte]int{}
	var long, short string
	if len(s) >= len(t) {
		long = s
		short = t
	} else {
		long = t
		short = s
	}
	for i := 0 ; i < len(long); i ++ {
		map1[long[i]]++
	}
	for i := 0 ; i < len(short); i++ {
		map2[short[i]]++
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}
// @lc code=end

