/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
package main
import "fmt"

func main() {
	var a []int
	a  = []int{2,2,3,2}	
	fmt.Println(singleNumber(a))

	a  = []int{0,1,0,1,0,1,99}	
	fmt.Println(singleNumber(a))
}

func singleNumber(nums []int) int {
	// dict  := map[int]int{}
	// for _, i := range nums {
	// 	dict[i]++
	// }
	// for k, v := range dict {
	// 	if v == 1 {
	// 		return k
	// 	}
	// }
	// return -1
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> uint(i)) & 1
		}
		result ^= (sum % 3) << uint(i)
	}
	return result

}
// @lc code=end

