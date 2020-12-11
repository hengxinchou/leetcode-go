/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start

// @lc code=end

package main
import "fmt"

func titleToNumber(s string) int {
	sum := 0
	for i := 0 ; i < len(s) ; i++ {
		sum = sum * 26 + int(s[i] - 'A' + 1)
	}
	return sum
}

func main() {
	var s string
	s = "A"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))

	s = "Z"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))

	s = "AA"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))

	s = "ZY"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))
	s = "AZ"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))
	s = "BA"
	fmt.Printf("%s, %d\n", s, titleToNumber(s))
}
