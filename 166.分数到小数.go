/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
package main

import "fmt"

func main() {

	var numerator, denominator int
	
	numerator = 1
	denominator = 2
	fmt.Printf("%d, %d, %f\n", numerator, denominator, fractionToDecimal(numerator, denominator))

	numerator = 2
	denominator = 1
	fmt.Printf("%d, %d, %f\n", numerator, denominator, fractionToDecimal(numerator, denominator))

	numerator = 2
	denominator = 3
	fmt.Printf("%d, %d, %f\n", numerator, denominator, fractionToDecimal(numerator, denominator))

	numerator = 4
	denominator = 333
	fmt.Printf("%d, %d, %f\n", numerator, denominator, fractionToDecimal(numerator, denominator))

	numerator = 1
	denominator = 5
	fmt.Printf("%d, %d, %f\n", numerator, denominator, fractionToDecimal(numerator, denominator))
}

func fractionToDecimal(numerator int, denominator int) string {

}

func ABS(v int ) int {
	if v < 0 {return -v}
	return v
}

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {return "NAN"}
	var tmp string
	if numerator*denominator < 0 {
		tmp += "-"
	}
	numerator,denominator = ABS(numerator),ABS(denominator)
	tmp += strconv.Itoa(numerator/denominator)
	num := numerator%denominator
	if num == 0 {return tmp}
	tmp += "."
	hm := make(map[int]int)
	rptPos := -1
	for {
		num *=10
		if pos,ok := hm[num]; ok {
			rptPos = pos
			break
		} else {
			hm[num] = len(tmp)
		}
		tmp += strconv.Itoa(num/denominator)
		num %= denominator
		if num == 0 {break}
	}
	if rptPos == -1 {
		return tmp
	}
	return fmt.Sprintf("%s(%s)",tmp[:rptPos], tmp[rptPos:])
}
// @lc code=end

