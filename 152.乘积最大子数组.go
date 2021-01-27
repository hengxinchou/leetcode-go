/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start

package main

import (
	"fmt"
	"math"
)

func main() {
	var a []int

	a = []int{2, 3, -2, 4}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{-2, 0, -1}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{0, -1, 4, -4, 5, -2, -1, -1, -2, -3, 0, -3, 0, 1, -1, -4, 4, 6, 2, 3, 0, -5, 2, 1, -4, -2, -1, 3, -4, -6, 0, 2, 2, -1, -5, 1, 1, 5, -6, 2, 1, -3, -6, -6, -3, 4, 0, -2, 0, 2}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{0, -1, 4, -4, 5, -2, -1, -1, -2, -3, 0, -3, -4, 4, 6}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{2, 2, -1, -5, 1, 1, 5, -6, 2, 1, -3, -6, -6, -3, 4}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{0}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{}
	fmt.Printf("%v, %d\n", a, maxProduct(a))

	a = []int{-2, 0}
	fmt.Printf("%v, %d\n", a, maxProduct(a))
}

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxF = max([]int{nums[i] * maxF, nums[i] * minF, nums[i]})
		minF = min([]int{nums[i] * maxF, nums[i] * minF, nums[i]})
		ans = max([]int{maxF, minF})
	}
	return ans
}

func max(a []int) int {
	maxV := math.MinInt32
	for i := 0; i < len(a); i++ {
		if a[i] > maxV {
			maxV = a[i]
		}
	}
	return maxV
}

func min(a []int) int {
	minV := math.MaxInt32
	for i := 0; i < len(a); i++ {
		if a[i] < minV {
			minV = a[i]
		}
	}
	return minV
}

func maxProduct1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := [][]int{}
	tmp := []int{}
	zeroFlag := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if len(tmp) > 0 {
				res = append(res, tmp)
			}
			tmp = []int{}
			zeroFlag = true
		} else {
			tmp = append(tmp, nums[i])
		}
	}
	if len(tmp) > 0 {
		res = append(res, tmp)
	}

	if zeroFlag {
		res = append(res, []int{0})
	}

	tmp = []int{}
	// fmt.Printf("res is %v\n", res)

	for i := 0; i < len(res); i++ {
		tmp = append(tmp, maxProduct2(res[i]))
	}

	maxV := math.MinInt32
	for _, v := range tmp {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func maxProduct2(nums []int) int {
	maxV := math.MinInt32

	n := len(nums)

	var dfs func(int, int)
	dfs = func(i, accu int) {
		if i == n {
			return
		}
		accu = accu * nums[i]

		if accu > maxV {
			maxV = accu
		}
		dfs(i+1, 1)
		dfs(i+1, accu)
	}

	dfs(0, 1)
	return maxV
}

// @lc code=end
