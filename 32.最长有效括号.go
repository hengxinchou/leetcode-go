/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
package main 
import "fmt"
func longestValidParentheses(s string) int {
	// stack := []byte{}
	// maxCount := 0
	// count := 0
	// for i := 0 ; i < len(s) ; i++ {
	// 	if s[i] == '(' {
	// 		stack = append(stack , s[i])
	// 		// fmt.Printf("first i is %d, s[i] is %s, stack are %s, maxCount is %d, count is %d\n",i, string(s[i]), string(stack), maxCount, count)
	// 	} else if len(stack) == 0 && s[i] == ')' {
	// 		maxCount = max(maxCount, count)
	// 		count = 0
	// 		// fmt.Printf("second, i is %d, s[i] is %s, stack are %s, maxCount is %d, count is %d\n",i, string(s[i]), string(stack), maxCount, count)
	// 	} else if s[i] == ')' {
	// 		prev := stack[len(stack)-1]
	// 		if prev == '(' {
	// 			stack = stack[:len(stack)-1]	
	// 		} 
	// 		if len(stack) == 0 {
	// 			count +=2
	// 		} else {

	// 		}

	// 		// fmt.Printf("3rd, i is %d, s[i] is %s, stack are %s, maxCount is %d, count is %d\n",i, string(s[i]), string(stack), maxCount, count)
	// 	}
	// }
	// return max(maxCount,count)


	maxAns := 0
    stack := []int{}
    stack = append(stack, -1)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                maxAns = max(maxAns, i - stack[len(stack)-1])
            }
        }
    }
    return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

func main (){
	a := ")()())"
	fmt.Printf("%s, %d\n", a, longestValidParentheses(a))

	a = "()(())(()"
	fmt.Printf("%s, %d\n", a, longestValidParentheses(a))
}

