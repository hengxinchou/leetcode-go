/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
package main
import "fmt"
func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}


	// if haystack == "" || needle == "" {
	// 	return -1
	// }



	size := len(haystack)
	size2 := len(needle)
	for i := 0 ; i < size - size2 + 1; i++ {
		var j int
		for j = 0 ; j < size2 ; j++{
			if haystack[i+j] == needle[j] {
				continue
			} else {
				break
			}
		}
		if j == size2 {
			return i
		}
	}
	return -1
}
// @lc code=end

func main(){
	a, b := "hello", "ll"
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "", "ll"
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "helllo", ""
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "helllo", "llo"
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "helllo", "he"
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "helllo", "he11111"
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))

	a, b = "", ""
	fmt.Printf("%s, %s, %d\n", a, b, strStr(a,b))
}

