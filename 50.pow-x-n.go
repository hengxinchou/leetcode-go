/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
package main

import "fmt"

func main() {
	
}
func myPow(x float64, n int) float64 {
	switch {
	case n == 0 || x == 1:
		return 1
	case n == 1:
		return x
	case x == 0：
		return 0
	case 
	}

	if x > 1 && n > 1 {
		sum := float64(1) 
		for i := 0 ; i < n; i++ {
			sum = sum * x
		}
		return sum
	}
	if x == 1 || n == 0 {
		return 0
	}
	}
}

func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickMul(x, n)
    }
    return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    y := quickMul(x, n/2)
    if n%2 == 0 {
        return y * y
    }
    return y * y * x
}
// @lc code=end

