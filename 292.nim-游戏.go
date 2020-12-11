/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start

// @lc code=end

package main
import "fmt"
func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}
	if n == 4  {
		return false
	}
	return canWinNim(n - 4)
}

func main() {
	// a := 6
	// fmt.Printf("%d, %t\n", a , canWinNim(a))
	// a := 7
	// fmt.Printf("%d, %t\n", a , canWinNim(a))
	for i := 0 ; i < 20 ; i++{
		fmt.Printf("%d, %t\n", i, canWinNim(i))
	}
}

