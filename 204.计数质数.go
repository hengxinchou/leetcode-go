/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
package main

import "fmt"
import "math"

type A struct {
	b map[int]int
} 

func isPrime(n int) bool {

	limit := int(math.Sqrt(float64(n)))
	for i := 2 ; i <= limit ; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func (a A) countPrimes(n int) int {
	var d int
	if  c, ok := a.b[n-1]; ok {
		d = c
	} else {
		d = a.countPrimes(n-1) 
	}

	if isPrime(n-1) {
		fmt.Printf("%d is prime\n", n -1)
		d += 1
	}
	return d
}

// func countPrimes(n int) int {
// 	signs := make([]bool, n)
// 	count := 0
// 	res := []int{}
// 	for i := 2 ; i < n ; i++ {
// 		if signs[i] {
// 			continue
// 		}
// 		count++
// 		res = append(res, i)
// 		for j := i + i ; j < n ; j = j+i{
// 			signs[j] = true
// 		}
// 	}
// 	fmt.Println(res)
// 	return count
// }
// @lc code=end




func main(){
	// fmt.Println(countPrimes(2))
	// fmt.Println(countPrimes(0))
	// fmt.Println(countPrimes(100000))
	// fmt.Println(countPrimes(1))

	a := A{}
	a.b = map[int]int{3: 1, 5: 2, 7: 3, 8: 4}
	c := a.countPrimes(19)
	fmt.Println(c)
}
