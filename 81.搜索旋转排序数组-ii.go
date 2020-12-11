/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
package main
import "fmt"
// import "sort"
func main() {
	var a []int 
	var target int 

	a = []int{2,5,6,0,0,1,2}
	target = 3
	fmt.Printf("array is %v, target %d \n", a, target)
	fmt.Printf("result: %t\n", search(a, target))

	a = []int{2,5,6,0,0,1,2}
	target = 3
	fmt.Printf("array is %v, target %d \n", a, target)
	fmt.Printf("result: %t\n", search(a, target))

	a = []int{1,3,1,1,1}
	target = 3
	fmt.Printf("array is %v, target %d \n", a, target)
	fmt.Printf("result: %t\n", search(a, target))
}

func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	var mid int 
	for start <= end {
		mid = start + (end - start) >> 2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] {
			start++
			continue
		}
		//前半部分有序
		if nums[start] < nums[mid] {
			//target在前半部分
			if nums[mid] > target && nums[start] <= target {
				end = mid - 1
			} else {  //否则，去后半部分找
				start = mid + 1
			}
		} else {
			//后半部分有序
			//target在后半部分
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {  //否则，去前半部分找
				end = mid - 1
			}
		}
	}
	//一直没找到，返回false
	return false
}
// @lc code=end

