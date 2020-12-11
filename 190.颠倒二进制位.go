/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start

// @lc code=end

package main

import "fmt"

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
    power := uint32(31)
    for num != 0 {
        ret += (num & 1) << power
        num = num >> 1
        power -= 1
    }
    return ret
}

func main() {

	var a = uint32(8)
	fmt.Printf("%d, %d\n", a, reverseBits(a))
}
