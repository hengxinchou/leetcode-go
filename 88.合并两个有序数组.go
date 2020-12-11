/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
package main 
import "fmt"
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// i , j := 0, 0 
	// end := m
	// for i < end && j < n {
	// 	// fmt.Printf("i is %d, j is %d, nums1 are %v, nums2 are %v\n", i, j, nums1, nums2)
	// 	if nums1[i] <= nums2[j] {
	// 		i++
	// 	} else {
	// 		end++
	// 		for k := end - 1 ; k > i ; k-- {
	// 			nums1[k] = nums1[k-1]
	// 		}
	// 		nums1[i] = nums2[j]
	// 		fmt.Printf("nums1 are %d\n", nums1)
	// 		i++
	// 		j++
	// 	}
	// }
	// // fmt.Printf("i is %d, j is %d, nums1 are %v\n", i,  j, nums1)
	// for j < n {
	// 	nums1[i] = nums2[j]
	// 	i++
	// 	j++
	// }
	// fmt.Printf("nums1 are %v\n", nums1)


	left, right, pivot := m - 1, n -1, m + n - 1
	for left >= 0 && right >= 0 {
		if nums1[left] >= nums2[right] {
			nums1[pivot] = nums1[left]
			pivot--
			left--
		} else {
			nums1[pivot] = nums2[right]
			pivot--
			right--
		}
	}
	for right >= 0 {
		nums1[pivot] = nums2[right]
		pivot--
		right--
	}
}
// @lc code=end\

func main () {
	nums1 := []int{1,2,3, 7, 0, 0,0}
	nums2 := []int{2,5,6}
	fmt.Printf("%v, %v\n", nums1, nums2)
	merge(nums1, 4, nums2, 3)
	fmt.Printf("%v, %v\n", nums1, nums2)
}
