/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
package main

import "fmt"

func main() {
  var coins []int	
  var amount int

  coins = []int{1,2,5}
  amount = 11
  fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))

  coins = []int{2}
  amount = 3
  fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))

  coins = []int{1}
  amount = 0
  fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))

  coins = []int{1}
  amount = 1
  fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))

  coins = []int{1}
  amount = 2
  fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	    max := amount + 1	
		dp := make([]int, amount + 1)

		for i := 0 ; i < amount + 1 ; i++ {
			dp[i] = max
		}

        dp[0] = 0
        for i := 1; i <= amount; i++ {
            for j := 0; j < len(coins); j++ {
                if coins[j] <= i {
                    dp[i] = min(dp[i], dp[i - coins[j]] + 1)
                }
            }
        }
        if dp[amount] > amount {
			return -1
		}  else {
			return dp[amount]
		}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
// @lc code=end

