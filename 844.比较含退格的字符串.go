/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
package main

import "fmt"
import "strings"

func backspaceCompare(S string, T string) bool {
	a := getResult(S)
	b := getResult(T)
	return a == b
}

func getResult(s string) string {
	stack1 := []byte{}
	for i := len(s) -1 ; i >= 0 ; i--{
		if len(stack1) == 0 {
			stack1 = append(stack1, s[i])
		} else {
			prev := stack1[len(stack1) -1 ]
			if s[i] != '#' && prev == '#' {
				stack1 = stack1[0:len(stack1) - 1]	
			} else {
				stack1 = append(stack1, s[i])
			}
		}
	}

	// fmt.Printf("first, %s, stack is %s\n", s, stack1)
	var result1 strings.Builder 
	for i := len(stack1) - 1 ; i >= 0; i-- {
		if stack1[i] != '#' {
			result1.WriteByte(stack1[i])
		}	
	}

	// fmt.Printf("second %s, result is %s\n", s, result1)
	return result1.String()
}


func main() {
	var a, b string
	a = "ab#c"
	b = "ad#c"
	fmt.Printf("%s, %s, %t\n", a, b, backspaceCompare(a,b))

	a = "ab##"
	b = "c#d#"
	fmt.Printf("%s, %s, %t\n", a, b, backspaceCompare(a,b))

	a = "a#c#"
	b = "b"
	fmt.Printf("%s, %s, %t\n", a, b, backspaceCompare(a,b))


	a = "a##c"
	b = "#a#c"
	fmt.Printf("%s, %s, %t\n", a, b, backspaceCompare(a,b))

	a = "nzp#o#g"
	b = "b#nzp#o#g"

	fmt.Printf("%s, %s, %t\n", a, b, backspaceCompare(a,b))
}

// @lc code=end

