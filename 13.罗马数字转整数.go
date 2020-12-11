/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
package main
import "fmt"
func romanToInt(s string) int {

	sum := 0
	n := len(s)
	for i :=  n- 1; i >= 0 ; i-- {
		switch string(s[i]) {
		case "I":
			if i < n -1 && (string(s[i+1]) == "V" || string(s[i+1]) == "X" ){
				sum -= 1
			} else {
				sum += 1
			}  
		case "V":
			sum += 5
		case "X":
			if i < n -1 && (string(s[i+1]) == "L" || string(s[i+1]) == "C" ){
				sum -= 10
			} else {
				sum += 10
			}
		case "L":
			sum += 50
		case "C":
			if i < n -1 && (string(s[i+1]) == "D" || string(s[i+1]) == "M" ){
				sum -= 100
			} else {
				sum += 100
			}
		case "D":
			sum += 500
		case "M":
			sum += 1000
		}
		fmt.Printf("sum is %s, %d\n", s[i:], sum)
	}
	return sum
}
// @lc code=end

func main (){
	// fmt.Printf("%s, %d\n", "XXVII", romanToInt("XXVII"))
	// fmt.Printf("%s, %d\n", "IV", romanToInt("IV"))
	// fmt.Printf("%s, %d\n", "LVIII", romanToInt("LVIII"))
	// fmt.Printf("%s, %d\n", "MCMXCIV", romanToInt("MCMXCIV"))
	fmt.Printf("%s, %d\n", "MDCCCLXXXIV", romanToInt("MDCCCLXXXIV")) //1884
}

