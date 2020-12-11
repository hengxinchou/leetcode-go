/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start

package main 
import "fmt"
// import "math"

func maxProfit(prices []int) int {
	profit := 0 

	for i := 1 ; i < len(prices) ; i++{
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}  
	}
	return profit
}
// @lc code=end

func main(){
	nums := []int{7,1,5,3,6,4}
	fmt.Printf("%v, %d\n", nums, maxProfit(nums))

	nums = []int{1,2,3,4,5}
	fmt.Printf("%v, %d\n", nums, maxProfit(nums))

	nums = []int{7,6,4,3,1}
	fmt.Printf("%v, %d\n", nums, maxProfit(nums))

	nums = []int{2,4,1}
	fmt.Printf("%v, %d\n", nums, maxProfit(nums))

	nums = []int{2,1,2,0,1}
	fmt.Printf("%v, %d\n", nums, maxProfit(nums))
}
