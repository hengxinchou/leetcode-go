/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

package main
import "fmt"
func twoSum(nums []int, target int) []int {

	res := []int{}
	a := map[int]int{}
	for i, v := range nums {
		left := target-v
		if _, existed := a[left]; existed {
			res = append(res, i)
			res = append(res, a[left])
			break
		}
		a[v] = i
	}
	return res
}
// @lc code=end


func main() {
	a := []int{2, 7, 11, 15}
	fmt.Printf("%v, %v\n", a, twoSum(a, 9))

	b := []int{3,0,-2,-1,1,2}
	fmt.Printf("%v, %v\n", b, twoSum(b,0))

	c := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}

	fmt.Printf("%v, %v\n", c, twoSum(c,0 ))
}
