/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start

// @lc code=end

package main

import "fmt"
// import "strings"

func convertToTitle(n int) string {
	s := ""
	for n > 0 {
		left := (n -1 ) % 26
		s = string('A' + left) + s 
		n = (n -1) / 26
	}
	return s

}

func main() {
	fmt.Printf("%d, %s\n", 1, convertToTitle(1))
	fmt.Printf("%d, %s\n", 25, convertToTitle(25))
	fmt.Printf("%d, %s\n", 26, convertToTitle(26))
	fmt.Printf("%d, %s\n", 27, convertToTitle(27))

	fmt.Printf("%d, %s\n", 28, convertToTitle(28))
	fmt.Printf("%d, %s\n", 29, convertToTitle(29))
	fmt.Printf("%d, %s\n", 701, convertToTitle(701))
	fmt.Printf("%d, %s\n", 52, convertToTitle(52))
	fmt.Printf("%d, %s\n", 53, convertToTitle(53))
	// fmt.Printf("%s", string('A' + 25 - 1))
}