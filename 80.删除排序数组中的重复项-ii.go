/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
package main

import "fmt"

func main() {
	var a []int
	var b int

	a = []int{1,1,1,2,2,3}
	fmt.Printf("%v\n", a)
	b = removeDuplicates(a)
	fmt.Printf("result is %d, %v\n", b, a)

	a = []int{0,0,1,1,1,1,2,3,3}
	fmt.Printf("%v\n", a)
	b = removeDuplicates(a)
	fmt.Printf("result is %d, %v\n", b, a)
}

func removeDuplicates(nums []int) int {
        // Initialize the counter and the second pointer.
		j := 1
		count := 1
        
        // Start from the second element of the array and process
        // elements one by one
        for i := 1; i < len(nums); i++ {
            // If the current element is a duplicate, increment the count.
            if (nums[i] == nums[i - 1]) {
                count++
            } else {
                // Reset the count since we encountered a different element
                // than the previous one.
                count = 1
            }
            // For a count <= 2, we copy the element over thus
            // overwriting the element at index "j" in the array
            if (count <= 2) {
				nums[j] = nums[i]
				j++
            }
        }
        return j
}

// @lc code=end

