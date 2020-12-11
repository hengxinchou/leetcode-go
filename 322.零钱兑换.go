/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func main() {
	var coins []int
	var amount int

	coins = []int{1, 2, 5}
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

	coins = []int{1, 2, 5}
	amount = 100
	fmt.Printf("%v, %d, %d\n", coins, amount, coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	map1 := map[int]int{}

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if v, ok := map1[amount]; ok {
			return v
		}
		if amount == 0 {
			return 0
		} else if amount < 0 {
			return -1
		}

		minV := math.MaxInt32
		for _, v := range coins {
			candiate := dfs(amount - v)
			if candiate < minV && candiate >= 0 {
				minV = candiate
			}
		}

		var res int
		if minV != math.MaxInt32 {
			res = minV + 1
		} else {
			res = -1
		}
		map1[amount] = res
		return res

	}
	a := dfs(amount)
	return a
}

// func coinChange1(coins []int, amount int) int {
// 	  // fmt.Printf("amount is %d\n", amount)
// 		if amount == 0 {
// 			return 0
// 		} else if amount < 0 {
// 			return -1
// 		}

// 		candiates := []int{}
// 		// fmt.Printf("2 amount is %d\n", amount)
// 		for _, v := range coins {
// 			  tmp := coinChange1(coins,  amount -v)
// 				candiates = append(candiates, tmp)
// 		}
// 		// fmt.Printf("candiates are %v\n", candiates)
// 		minV := math.MaxInt32
// 		for _, v := range candiates {
// 			if v < minV && v >= 0 {
// 				minV = v
// 			}
// 		}
// 		if minV != math.MaxInt32 {
// 			return minV + 1
// 		} else {
// 			return -1
// 		}
// }

// func coinChange3(coins []int, amount int) int {
// 		sort.Slice(coins, func(i, j int) bool {
// 			return coins[i] > coins[j]
// 		})
// 		fmt.Printf("coins are %v\n",coins)
// 		count := 0
// 		for amount > 0 {
// 			fmt.Printf("amount is %d\n", amount)
// 			for _, v := range coins {
// 				if v <= amount {
// 					amount -= v
// 					count++
// 					break
// 				}
// 			}
// 		}
// 		if amount < 0 {
// 			return -1
// 		}
// 		return count
// }

// func coinChange2(coins []int, amount int) int {

// 	var dfs func(int)

// 	count := 0
// 	dfs = func(left int){
// 		if 0 == left  {
// 			count++
// 			return
// 		}
// 		if left < 0 {
// 			return
// 		}

// 		for _, v := range coins {
// 			dfs(left-v)
// 		}
// 	}
// 	dfs(amount)
// 	return count
// }
// @lc code=end

//动态规划（压缩空间）
func coinChange4(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始值
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
