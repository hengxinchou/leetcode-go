/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
package main
import "fmt"
func main() {
	
	fmt.Printf("%d, %t\n", 1, isPowerOfTwo(1))
	fmt.Printf("%d, %t\n", 16, isPowerOfTwo(16))
	fmt.Printf("%d, %t\n", 4, isPowerOfTwo(4))
	fmt.Printf("%d, %t\n", 5, isPowerOfTwo(5))
	fmt.Printf("%d, %t\n", 0, isPowerOfTwo(0))
	fmt.Printf("%d, %t\n", -2, isPowerOfTwo(-2))
	fmt.Printf("%d, %t\n", -1, isPowerOfTwo(-1))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n & (n - 1) == 0 {
		return true
	} else {
		return false
	}
}
// @lc code=end

