/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
package main 
import "fmt"
import "strconv"

func addBinary(a string, b string) string {
	sizeA := len(a)
	sizeB := len(b)
	var maxSize int
	if sizeA > sizeB {
		maxSize = sizeA
	} else {
		maxSize = sizeB
	}
	var sum, left int
	carryOn := 0
	res := ""
	for i := 0 ; i < maxSize ; i++ {
		sum = 0
		if i < sizeA {
			sum += int(a[sizeA - i - 1] - '0')
		} 
		if i < sizeB {
			sum += int(b[sizeB - i -1] - '0')
		}
		sum += carryOn
		left = sum % 2 
		carryOn = sum / 2
		res = strconv.Itoa(left) + res
	}
	if carryOn == 1 {
		res = "1" + res
	}
	return res
}
// @lc code=end

func main (){
	fmt.Println(addBinary("1010","1011"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("100", "110010"))
}
