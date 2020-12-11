/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package main
import "fmt"
import "math"
// import "strings"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	size := len(s)
	i := 0
	for i < size && s[i] == ' ' {
		i++
	}
	if i == size {
		return 0
	}
	s = s[i:]
 
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	n := 0
	// s2 := []byte(s)
	for  _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)

		if n > math.MaxInt32 {
			if s0[0] == '-' {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	// fmt.Printf("n is %d\n", n)

	if s0[0] == '-' {
		n = -n
	}
	return n


	// s = strings.TrimSpace(s)
	// result := 0
	// sign := 1

	// for i, v := range s {
	// 	if v >= '0' && v <= '9' {
	// 		result = result*10 + int(v-'0')
	// 	} else if v == '-' && i == 0 {
	// 		sign = -1
	// 	} else if v == '+' && i == 0 {
	// 		sign = 1
	// 	} else {
	// 		break
	// 	}
	//    // 数值最大检测
	//    fmt.Printf("result is %d\n", result)
	// 	if result > math.MaxInt32 {
	// 		if sign == -1 {
	// 			return math.MinInt32
	// 		}
	// 		return math.MaxInt32
	// 	}
	// }

	// return sign * result

}
// @lc code=end

func main(){
	// fmt.Println(myAtoi("    -43"))
	// fmt.Println(myAtoi("42"))
	// fmt.Println(myAtoi("4193 with words"))
	// fmt.Println(myAtoi("words and 987"))
	// fmt.Println(myAtoi("-91283472332"))
	fmt.Println("9223372036854775808")
	fmt.Println(myAtoi("9223372036854775808"))
	// fmt.Println(myAtoi("21474836460"))
	// fmt.Println(myAtoi("0000000000012345678"))
	
}