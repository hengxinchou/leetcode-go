/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start

package main

import "fmt"
import "sort"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 1
	}

	sort.Ints(nums)

	size := len(nums)

	if nums[0] > 1 {
		return 1
	}

	if nums[size -1] <= 0 {
		return 1
	}

	var res int 
	if nums[0] == 1 {
		j := 1
		for i := 0 ; i < size ; i++ {
			// fmt.Printf("i is %d, nums[%d] is  %d\n", i, i, nums[i])
			if nums[i] == j {
				continue
			} else {
				j++
				if nums[i] == j {
					continue
				} else {
					return j
				}
			}
		}
		res = j + 1
	} else if nums[0] <= 0 {
		for i, v := range nums {
			if v > 0 {
				res = firstMissingPositive(nums[i:])
				break
			}
		}
		
	} 
	return res
}
// @lc code=end

func main (){
	a1 := []int{1,2,0}
	fmt.Printf("%v, %d\n", a1, firstMissingPositive(a1))

	a2 := []int{3,4,-1,1}
	fmt.Printf("%v, %d\n", a2, firstMissingPositive(a2))

	a3 := []int{7,8,9,11,12}
	fmt.Printf("%v, %d\n", a3, firstMissingPositive(a3))

	a4 := []int{7,8,9,11,12, 1}
	fmt.Printf("%v, %d\n", a4, firstMissingPositive(a4))

	a5 := []int{}
	fmt.Printf("%v, %d\n", a5, firstMissingPositive(a5))

	a6 := []int{0}
	fmt.Printf("%v, %d\n", a6, firstMissingPositive(a6))

	a7 := []int{-1}
	fmt.Printf("%v, %d\n", a7, firstMissingPositive(a7))

	a8 := []int{1}
	fmt.Printf("%v, %d\n", a8, firstMissingPositive(a8))

	a9 := []int{0,2,2,1,1}
	fmt.Printf("%v, %d\n", a9, firstMissingPositive(a9))
}
