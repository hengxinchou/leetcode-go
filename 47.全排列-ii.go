/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
package main

import "fmt"


func main() {
	var a []int
	// a = []int{1,1,2}
	
	// fmt.Printf("%v\n", permuteUnique(a))

	a = []int{2,2,1,1}
	fmt.Printf("%v\n", permuteUnique(a))
}

func permuteUnique(nums []int) [][]int{
	if len(nums) == 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		if nums[0] != nums[1] {
			return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
		} else {
			return [][]int{nums}
		}
		
	}

	res := [][]int{}
	isDone := map[int]bool{}
	for i := 0 ; i < len(nums) ; i++ {
		// if i != 0 && nums[i] == nums[0] {
		// 	continue
		// }
		if _, ok := isDone[nums[i]] ; ok {
			continue 
		} else {
			isDone[nums[i]] = true
		}

		newA := make([]int, len(nums))
		copy(newA, nums)
		newA[0], newA[i] = newA[i], newA[0]
		
		res2 := permuteUnique(newA[1:])
		fmt.Printf("one round, newA[0] is %d, ", newA[0])
		for _, tmp := range res2 {
			tmp2 := append([]int{newA[0]}, tmp...)
			fmt.Printf("%v, ", tmp2)
			res = append(res, tmp2)
		}
		fmt.Println()
	}
	return res
}
// @lc code=end

