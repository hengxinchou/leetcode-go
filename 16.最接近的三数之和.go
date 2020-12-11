/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
package main

import "fmt"
import "math"

func main() {
	var a []int
	var target int
	
	// a = []int{-1,2,1,-4}
	// target = 1 
	// fmt.Printf("%v, %d, %d\n", a, target, threeSumClosest(a, target))

	a = []int{1,2,4,8,16,32,64,128}
	target = 82
	fmt.Printf("%v, %d, %d\n", a, target, threeSumClosest(a, target))
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	min := math.MaxInt64
	var mostApproach int
	for first := 0 ; first < n - 2 ; first++ {
		for second := first + 1 ; second < n - 1; second++ {
			third := n - 1
			fmt.Printf("11111: %d, %d, %d\n", first, second, third)
			for third > second {
				sum := nums[third] + nums[first] + nums[second]
				distance := getDistance(sum, target)
				fmt.Printf("%d, %d, %d, sum is %d, target is %d, distance is %d \n", first, second ,third, sum, target, distance)
				if distance == 0 {
					fmt.Printf("1: %d, %d, %d\n", first, second, third)
					return sum
				} else if distance < min {
					fmt.Printf("2: %d, %d, %d, min is %d\n", first, second, third, distance)
					min = distance
					mostApproach = sum
				}
				third--
			}
		}
	} 

    return mostApproach
}
// @lc code=end

func getDistance(a, b int) int {
	return abs(a - b )
}

func abs(a int) int {
	if a < 0 {
		return (-1) * a
	} else {
		return a
	}
}
