/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	mapA := map[int]int{3: 1, 5: 2, 7: 3, 8: 4}

	isPrime := func(n int) bool {
		limit := int(math.Sqrt(float64(n)))
		for i := 2; i <= limit; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	var result int
	if v, ok := mapA[n-1]; ok {
		result = v
	} else {
		result = countPrimes(n - 1)
	}

	if isPrime(n - 1) {
		// fmt.Printf("%d is prime\n", n-1)
		result += 1
	}
	return result
}

func countPrimes1(n int) int {
	signs := make([]bool, n)
	count := 0
	res := []int{}
	for i := 2; i < n; i++ {
		// 如果 i是非质数则跳过
		if signs[i] {
			continue
		}
		count++              // 是质数
		res = append(res, i) // 把质数都添加到result
		// 倍数都不是质数
		for j := i + i; j < n; j = j + i {
			signs[j] = true
		}
	}
	// fmt.Println(res)
	return count
}

func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime { // 默认都是质数
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

// @lc code=end

func main() {
	var a int
	a = 19
	fmt.Printf("%d, %d\n", a, countPrimes1(a))

	a = 5
	fmt.Printf("%d, %d\n", a, countPrimes1(a))

	a = 100000
	fmt.Printf("%d, %d\n", a, countPrimes(a))
}
