/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
package main
import "fmt"
import "sort"
func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	for k, i := range A {
		res[k] = i * i
	}
	sort.Ints(res)
	return res
}
// @lc code=end



func main() {
	a := []int{-4,-1,0,3,10}
	fmt.Printf("%v, %v\n", a, sortedSquares(a))

	a = []int{-7,-3,2,3,11}
	fmt.Printf("%v, %v\n", a, sortedSquares(a))
}

