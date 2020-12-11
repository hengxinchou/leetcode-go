/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package main
import "fmt"
func main() {
	var a []int

	a = []int{4,5,6,7,0,1,2}
	fmt.Printf("%v,%d\n", a, search(a, 0))

	a = []int{4,5,6,7,0,1,2}
	fmt.Printf("%v,%d\n", a, search(a, 3))


	a = []int{5,1,2}
	fmt.Printf("%v,%d\n", a, search(a, 5))

	a = []int{5,1,2}
	fmt.Printf("%v,%d\n", a, search(a, 2))
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	down := 0
	upper := len(nums) -1
	var mid int

	for  down <= upper {
		mid = down+((upper-down)>>1)
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			// 赶 down upper到有序数组中查找
			if nums[0] <= target &&  target < nums[mid]  {
				upper = mid - 1
			} else {
				down = mid + 1 
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums) -1] {
				down = mid + 1
			} else {
				upper = mid - 1
			}
		}

	}
	return -1
}

// @lc code=end

