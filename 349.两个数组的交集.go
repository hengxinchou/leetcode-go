/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
package main
import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	a := map[int]bool{}
	b := map[int]bool{}

	for _, v := range nums1 {
		a[v] = true
	}

	res := []int{}

	for _, v := range nums2 {
		if _, ok := a[v]; ok {
			b[v] = true
		}
	}

	for k, _ := range b {
		res = append(res, k)
	}
	return res
}
// @lc code=end
func main (){
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}

	fmt.Printf("%v, %v, intersection is %v\n", nums1, nums2, intersection(nums1, nums2))

	nums11 := []int{4,9,5}
	nums22 := []int{9,4,9,8,4}
	fmt.Printf("%v, %v, intersection is %v\n", nums11, nums22, intersection(nums11, nums22))
}