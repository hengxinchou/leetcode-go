/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start

// @lc code=end

package main
import "fmt"
func hammingWeight(num uint32) int {
	sum := 0
    for num != 0 {
        sum++
        num =  num & (num - 1)
    }
    return sum
}

func main() {
	a := uint32(7)
	fmt.Printf("%d, %d\n", a, hammingWeight(a))
}