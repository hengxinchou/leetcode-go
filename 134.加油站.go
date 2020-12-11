/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
package main

import "fmt"
import "math"

func main() {

	var gas, cost []int
	gas  = []int{1,2,3,4,5}
	cost = []int{3,4,5,1,2}
	fmt.Println(canCompleteCircuit(gas, cost))

	gas  = []int{2,3,4}
	cost = []int{3,4,3}
	fmt.Println(canCompleteCircuit(gas, cost))

	gas = []int{5,8,2,8}
	cost = []int{6,5,6,6}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit2(gas []int, cost []int) int {
	tmp := make([]int, len(gas))

	for i := 0 ; i < len(gas) ; i++ {
		tmp[i] = gas[i] - cost[i]
	}

	index := findBigger(tmp)

	newTmp := tmp[index:len(tmp)]
	newTmp = append(newTmp, tmp[0:index]...)
	sum := 0
	for i := 0 ; i < len(newTmp) ; i++ {
		sum += newTmp[i]	
		if sum < 0 {
			return -1
		}
	}
	return index
}

func findBigger(a []int) int {
	maxV := math.MinInt32
	var maxI int
	for i := 0 ; i < len(a) ; i ++ {
		if a[i] > maxV {
			maxV = a[i]
			maxI = i
		}
	}
	return maxI
}
// @lc code=end

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	for i := 0 ; i < n ; {
		count := 0 
		sumOfGas := 0
		sumOfCost := 0 
		for count < n {
			j := (i + count ) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfGas < sumOfCost {
				break
			}
			count++
		}
		if count == n {
			return i
		} else {
			i += count + 1
		}
	}
	return -1
}