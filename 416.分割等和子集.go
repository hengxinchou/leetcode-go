/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
package main 
import "fmt"
import "sort"
func canPartition(nums []int) bool {
	sum := 0 
	for _, v := range nums {
		sum += v
	}
	if sum % 2 != 0 {
		return false
	}

	target := sum / 2 
	sort.Ints(nums)

	sum2 := 0
	for _, v := range nums {
		sum2 += v 
		if sum2 == target {
			return true
		}
		if sum2 > target {
			return false
		}
	}
	return false
}
// @lc code=end

func main(){
	a := []int{1, 5, 11, 5}
	fmt.Printf("%v, %t\n", a, canPartition(a))

	a = []int{1, 2, 3, 5}
	fmt.Printf("%v, %t\n", a, canPartition(a))

	a = []int{2,2,1,1}
	fmt.Printf("%v, %t\n", a, canPartition(a))
}