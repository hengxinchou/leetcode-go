/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

package main
import "fmt"

func sum2(s []int) int{
	sum1 := 0 
	for _, v := range s {
		sum1 += v
	}
	return sum1
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := 0
	size := len(nums)
	result := 0
	sum := nums[0]
	for {
		// 总数比 s 大，则i向右前进
		if sum >= s {
			if i >= size {
				break
			}

			leng := j - i + 1
			// if leng == 1 {
				// result = j - i + 1
				// break
			// }
			if leng < result || result == 0 {
				result = leng
			}
			fmt.Printf("first, i is %d, j is %d, sum is %d\n", i, j, sum2(nums[i:j+1]))
			if i <= j {
				sum -= nums[i]
				i++
			}

		// 总数比 s 小，则i向右前进
		} else {
			fmt.Printf("second: i is %d, j is %d, sum is %d\n", i, j, sum2(nums[i:j+1]))

			j++ 
			if j >= size {
				break
			} 
		
			sum += nums[j]
		} 
	}
	return result
}
// @lc code=end


func main() {
	s := 7
	nums := []int{2,3,1,2,4,3}
	fmt.Printf("%v\n", nums)
	fmt.Printf("min sub is %d\n", minSubArrayLen(s, nums))
	// b := []int{}
	// b = append(b, nums[5:7]...)
	// fmt.Println(b)
}
