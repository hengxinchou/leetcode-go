/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
package main
import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}

	for first := 0 ; first < n - 3 ; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1 ; second < n - 2 ; second++{
			if second > first + 1  && nums[second] == nums[second - 1] {
				continue
			}
			for third := second + 1 ; third < n -1 ; third++ {
				if third > second + 1  && nums[third] == nums[third - 1] {
					continue
				}
				
				left := target - nums[first] - nums[second] - nums[third]
				fourth := n - 1
				for  nums[fourth] > left && fourth > third  {
					fourth--
				}
				if fourth == third {
					continue
				}
				if nums[fourth] + nums[first] + nums[second] + nums[third] == target {
					res = append(res, []int{nums[first],nums[second], nums[third], nums[fourth]})
				}
			}
		}
	}
	return res
}
// @lc code=end

func main() {
	a := []int{1, 0, -1, 0, -2, 2}
	fmt.Printf("original are %v, res is %v\n", a, fourSum(a, 0 ))

	b := []int{3,0,-2,-1,1,2}
	fmt.Printf("original are %v, res is %v\n", b, fourSum(b, 0))

	c := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}

	fmt.Printf("original are %v, res is %v\n", c, fourSum(c, 0))
}