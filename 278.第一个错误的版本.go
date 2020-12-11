/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

package main

import "fmt"



func isBadVersion(version int) bool {
	a := []bool{false, true, true}
	return a[version+1]
}

func getPrev(m int) int {
	if m < 1 {
		return 1
	} else {
		return m
	}
}
func firstBadVersion(n int) int {

	mid := (1 + n) / 2
	prev := getPrev(mid-1)
	low := 1 
	high := n
	for low <= high {

		if isBadVersion(mid) && (!isBadVersion(prev) || mid == 1) {
			return mid
		} 

		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (high + low) / 2 
		prev = getPrev(mid - 1)
	}
	return -1
}
// @lc code=end

func main(){
	fmt.Println(firstBadVersion(1))
}
