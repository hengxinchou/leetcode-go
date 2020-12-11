/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
package main 
import "fmt"
func twoSum(numbers []int, target int) []int {

	low, upper := 0, len(numbers) - 1 

	var sum int

	for low < upper {
		sum = numbers[low] + numbers[upper] 
		if sum == target {
			return []int{low+1, upper+1}
		} else if sum > target {
			upper--
		} else {
			low++
		}
	} 
	return []int{-1, -1}
}
// @lc code=end

func main(){
	a := []int{2, 7, 11, 15}
	target := 18
	fmt.Printf("%v, %v\n", a, twoSum(a, target))
}

