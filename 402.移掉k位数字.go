/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */
 package main

import "fmt"

// @lc code=start
// func removeKdigits(num string, k int) string {
// 	if k >= len(num) {
// 		return "0"
// 	}
// 	stack := make([]int, len(num))
// 	index := 0
// 	for j, v := range num {
// 		i, err := strconv.Atoi(string(v))
// 		if err != nil {
// 			panic("strconv.Atoi goes wrong")
// 		}
// 		// fmt.Printf("is is %d, index is %d, stack %v\n", i, index, stack)
// 		if j == 0 {
// 			stack = append(stack[:0], i)
// 			index++
// 			continue 
// 		}

// 		// stack.pop

// 		// fmt.Printf("i is is %d, index is %d, stack %v\n", i, index, stack)
// 		for index > 0 && k > 0 {
// 			val := stack[index-1]
// 			if i >= val {
// 				break
// 			}
// 			index--
// 			k--
// 			// fmt.Printf("for val is %d, index is %d, stack %v\n", val, index, stack)

// 		}

// 		// fmt.Printf("index is %d, stack %v\n", index, stack)
// 		if i == 0 && index == 0 {
			
// 			stack = append(stack[:index], 0)
// 		} else {
// 			stack = append(stack[0:index], i)
// 			index++
// 		}
			
// 		// fmt.Printf("final if, i is %d, index is %d, stack %v\n", i, index, stack)
// 	}

// 	for k > 0 {
// 		stack = stack[0:index-1]
// 		index--
// 		k--
// 	}

// 	s := ""
// 	for _, v := range stack {
// 		s = s + strconv.Itoa(v)
// 	}
// 	return s
// }
// @lc code=end

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	var stack []rune 
	var result string

	for _, j := range num {
		i := j - '0'
		fmt.Printf("%d\n", i)
		for len(stack) != 0 && stack[len(stack)-1] > i && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		} 

		// 预防 0 首先进入stack
		if i != 0 || len(stack) != 0 {
			stack = append(stack, i)
		}
	}

	for len(stack) != 0  && k > 0 {
		stack = stack[:len(stack) - 1]
		k--
	}



	for _, l := range stack {
		result += string('0' + l)
	}

	if result == "" {
		return "0"
	}

	return result
}

func main(){
	// fmt.Printf("1432219, 3, the result %s\n", removeKdigits("1432219", 3))
	// fmt.Printf("10200, 1, the result %s\n", removeKdigits("10200", 1))
	// fmt.Printf("112, 1, the result %s\n", removeKdigits("112", 1))
	// fmt.Printf("10, 1, the result %s\n", removeKdigits("10", 1))
	// fmt.Printf("1111111, 3, the result %s\n", removeKdigits("1111111", 3))
	fmt.Printf("1234567890, 9, the result %s\n", removeKdigits("1234567890", 9))
}