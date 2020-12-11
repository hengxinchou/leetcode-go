/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package main
import "fmt"
import "strings"

func main() {
	var s string 
	var res []string
	s = "29"
	res = letterCombinations(s)
	fmt.Printf("digits is %s, result is %v, len is %d\n", s, res, len(res))

	s = "239"
	res = letterCombinations(s)
	fmt.Printf("digits is %s, result is %v, len is %d\n", s, res, len(res))
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	mapPhones := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	collectStr := []string{}
	for i := 0 ; i < len(digits) ; i++ {
		tmp, _ := mapPhones[digits[i]]
		collectStr = append(collectStr, tmp)
	}

	fmt.Printf("collectStr are %v\n", collectStr)
	finalRes := strings.Split(collectStr[0], "")

	fmt.Printf("1: finalRes are %v\n", finalRes)
	collectStr = collectStr[1:]

	for len(collectStr) > 0 {
		tmpRes := []string{}
		tmpS := collectStr[0]
		for _, midiumS := range finalRes {
	  		for i := 0 ; i < len(tmpS) ; i++ {
				tmpRes = append(tmpRes, midiumS + string(tmpS[i]))
			}
		}
		collectStr = collectStr[1:]
		finalRes = tmpRes
	}
	return finalRes
}
// @lc code=end

