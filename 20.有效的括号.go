/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package main

import "fmt"
// import "container/list"

func main(){
	var s string 
	s = "(())"
	fmt.Printf("%s, %t\n", s, isValid(s))

	s = "(()))"
	fmt.Printf("%s, %t\n", s, isValid(s))

	s = "()[]{}"
	fmt.Printf("%s, %t\n", s, isValid(s))
}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				fmt.Printf("stack is %v\n", stack)
				return false
			}
			prev := stack[len(stack)-1]
			stack = stack[0: len(stack) -1]
			tmp := string(prev) + string(s[i])
			fmt.Printf("tmp is %s\n", tmp)
			if tmp == "()" || tmp == "[]" || tmp == "{}" {
			  continue
			} else {
				fmt.Printf("this end, tmp is %v\n", tmp)
				return false
			}
		}
	}
	fmt.Printf("this end is %v\n", stack)
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}
// @lc code=end
