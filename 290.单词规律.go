/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
package main
import "fmt"
import "strings"

func wordPattern(pattern string, s string) bool {
	
	str_arr := strings.Split(s, " ")
	if len(str_arr) != len(pattern) {
		return false
	}
	a := map[byte]string{}
	used := map[string]bool{}
	for i := 0 ; i < len(pattern) ; i++ {
		key := pattern[i]
		provideValue := str_arr[i]
		if requireValue, ok := a[key]; !ok {
			if _, ok = used[provideValue]; !ok {
				a[key] = provideValue
				used[provideValue] = true
			} else {
				return false
			}	
		} else if requireValue != provideValue {
			return false
		}
	}
	return true
}
// @lc code=end

func main(){
	fmt.Printf("abba, dog cat cat dog, is %t\n", wordPattern("abba", "dog cat cat dog"))
	fmt.Printf("abba, dog cat cat fish, is %t\n", wordPattern("abba", "dog cat cat fish"))
	fmt.Printf("aaaa, dog cat cat dog, is %t\n", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Printf("abba, dog dog dog dog, is %t\n", wordPattern("abba", "dog dog dog dog"))
}