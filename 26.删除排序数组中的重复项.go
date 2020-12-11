/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
package main
import "fmt"
func removeDuplicates(nums []int) int {
	// i := 1
	// for i < len(nums)  {
	// 	if nums[i] == nums[i-1] {
	// 		nums = append(nums[0:i], nums[i+1:]...)
			
	// 	} else {
	// 		i++
	// 	}
	// }
	// // fmt.Printf("result is %v\n", nums)
	// return len(nums) 

	prev := 0 
	for i := 1 ; i< len(nums) ; i++ {
		if nums[i] != nums[prev] { // pivot其实是前一个值，
			prev++
			nums[prev] = nums[i]
		}
	}
	return prev + 1
}
// @lc code=end

func main(){
	// nums := []int{1,1,2}
	// fmt.Printf("%v\n", nums)
	// fmt.Printf("%v, %v\n", nums, removeDuplicates(nums))
	// fmt.Printf("%v\n", nums)

	// nums = []int{0,0,1,1,1,2,2,3,3,4}
	nums := []int{0,0,1,2,3,4}
	// fmt.Printf("%v\n", nums)
	fmt.Printf("%v, %v\n", nums, removeDuplicates(nums))
	fmt.Printf("%v\n", nums)
}