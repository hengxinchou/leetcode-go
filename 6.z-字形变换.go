/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
package main

import "fmt"

func main() {
	// tmp := make([]byte, 2)
	// tmp[1]
	// fmt.Printf("%s\n", string(tmp))
	var tmp byte
	fmt.Printf("%b\n", tmp)
	var s string
	s = "LEETCODEISHIRING"

	fmt.Printf("%s, %s\n", s, convert(s, 3))

	s = "LEETCODEISHIRING"
	fmt.Printf("%s, %s\n", s, convert(s, 4))

	s = "AB"
	fmt.Printf("%s, %s\n", s, convert(s, 1))
}

// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s 
// 	}
// 	low := numRows - 1
// 	upper := 0

// 	res := [][]byte{}

// 	for i := 0 ; i < numRows ; i++ {
// 		tmp := make([]byte, len(s))
// 		res = append(res, tmp)
// 	}
// 	t := 0
// 	j := 0
// 	i := 0
// 	for t < len(s) {
// 		for i < low && t < len(s) {
// 			res[i][j] = s[t]
// 			i++
// 			t++ 
// 		}
// 		if t == len(s) {
// 			break
// 		}
// 		res[i][j] = s[t] // i is low
// 		j++
// 		i--
// 		t++
// 		for i > upper && t < len(s) {
// 			res[i][j] = s[t]
// 			t++
// 			i--
// 			j++
// 		}
// 	}

// 	res2 := ""
// 	for m := 0 ; m < numRows ; m++ {
// 		for n := 0 ; n < len(s) ; n++ {
// 			if res[m][n] != byte(0) {
// 				res2 += string(res[m][n])
// 			}
// 		}
// 	}
// 	return res2
// }
// @lc code=end

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    rows := make([]string, numRows)
    n := 2 * numRows - 2
    for i, char := range s {
        x := i % n
        rows[min(x, n - x)] += string(char)
    }
    return strings.Join(rows, "")
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}