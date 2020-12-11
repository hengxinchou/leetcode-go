/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
package main

import "fmt"
import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
    sort.Ints(nums)
	ans := make([][]int, 0)
	
	for first := 0 ; first < n - 2 ; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}
		
		for second := first + 1 ; second < n - 1; second++ {
			if second > first +1 && nums[second] == nums[second-1] {
				continue
			}
			target := -1 * (nums[first] + nums[second])
			third := n - 1
			for nums[third] > target && third > second {
				third--
			} 
			if third == second {
				break
			}

			if nums[third] + nums[first] + nums[second] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	} 

    return ans
}
// @lc code=end

func main() {
	a := []int{-1,0,1,2,-1,-4}
	fmt.Printf("%v, %v\n", a, threeSum(a))

	b := []int{3,0,-2,-1,1,2}
	fmt.Printf("%v, %v\n", b, threeSum(b))

	c := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}

	fmt.Printf("%v, %v\n", c, threeSum(c))
}
