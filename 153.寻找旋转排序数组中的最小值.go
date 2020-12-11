/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
package main
import "math"
import "fmt"

func main() {
	var a []int
	a = []int{4,5,6,7,0,1,2}
	fmt.Printf("%v, %d\n", a, findMin(a))

	a = []int{3,4,5,1,2}
	fmt.Printf("%v, %d\n", a, findMin(a))

	a = []int{1}
	fmt.Printf("%v, %d\n", a, findMin(a))
}

func findMin(nums []int) int {
	minV := math.MaxInt32
	for _, v := range nums {
		if v < minV {
			minV = v
		}
	}
	return minV
}
// @lc code=end

