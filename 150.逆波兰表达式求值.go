/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
// 前提是 s 是有效的表达式
package main
import "fmt"
import "strconv"

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"  
}

func doOperate(a, b int, ope string) (res int){
	switch ope {
	case "+":
		res =  a + b 
	case "-":
		res =  a - b
	case "*":
		res =  a * b
	case "/":
		res =  a / b 	
	}
	return 
}

func evalRPN(tokens []string) int {
	if tokens == nil {
		return 0
	}
	stack := []int{}

	var res int
	for i := 0 ; i < len(tokens) ; i++{
		if !isOperator(tokens[i]) {
			tmp, _ := strconv.Atoi(tokens[i])
			stack = append(stack, tmp)
		} else {
			a := stack[len(stack) - 2]
			b := stack[len(stack) - 1]
			stack = stack[0:len(stack)-2]
			fmt.Printf("i is %d, a is %d, b is %d, operator is %s, stack is %v\n", i, a, b, tokens[i], stack)
			res = doOperate(a, b, tokens[i])
			stack = append(stack, res)
		}
	}
	return stack[0]
}
// @lc code=end

func main(){
	a1 := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(a1))

	a2 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(a2))

	a3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(a3))

	a4:= []string{"18"}
	fmt.Println(evalRPN(a4))
}

