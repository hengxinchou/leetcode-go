/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
package main 
import "fmt"
func singleNumber(nums []int) int {
	a := map[int]int{}
	for _, v := range nums {
		a[v]++
	}
	for k ,v := range a {
		if v == 1 {
			return k
		}
	}
	return -1
}
// @lc code=end

func main(){
	a := []int{2,2,1}
	fmt.Printf("%v, %d\n", a, singleNumber(a))

	a = []int{4,1,2,1,2}
	fmt.Printf("%v, %d\n", a, singleNumber(a))
}
