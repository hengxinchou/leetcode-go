/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
package main
import "fmt"
import "unicode"
func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	i := 0 
	j := len(s) - 1
	ss := []rune(s)
	for i < j {
		for i < len(ss) && !unicode.IsDigit(ss[i]) && !unicode.IsLetter(ss[i]) {
			i++
		}
		if i == len(ss) {
			return true
		}
		for j >= 0 && !unicode.IsDigit(ss[j]) && !unicode.IsLetter(ss[j])  {
			j--
		}
		// fmt.Printf("first %s, %s\n", string(ss[i]), string(ss[j]))
		if  s[i] == s[j] || unicode.ToUpper(a) == unicode.ToUpper(b) {
			i++
			j--
			continue
		} else {
			// fmt.Printf("%s, %s\n", string(ss[i]), string(ss[j]))
			return false
		}
	}
	return true
}

func isSame(a, b rune) bool {
	return unicode.ToUpper(a) == unicode.ToUpper(b)
}
// @lc code=end

func main(){
	// fmt.Printf("%d\n", 'A' - 'a')
	var s string
	s = "A man, a plan, a canal: Panama"
	fmt.Printf("%s, %t\n", s, isPalindrome(s))

	s = "race a car"
	fmt.Printf("%s, %t\n", s, isPalindrome(s))

	s = ".,"
	fmt.Printf("%s, %t\n", s, isPalindrome(s))

	s = "0P"
	fmt.Printf("%s, %t\n", s, isPalindrome(s))
}

