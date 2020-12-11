/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
package main
import "math"
import "fmt"

func maxProfit(prices []int) int {

	minPrice := math.MaxInt64

	maxProfit := 0

	for i := 0 ; i < len(prices) ; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit

}
// @lc code=end


func main(){
	a := []int{7,1,5,3,6,4}
	fmt.Printf("%v, %d\n", a, maxProfit(a))

	a = []int{7,6,4,3,1}
	fmt.Printf("%v, %d\n", a, maxProfit(a))
}
