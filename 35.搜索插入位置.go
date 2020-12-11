/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var i int
	for i = 0 ; i < len(nums) ; i++ {
		if nums[i] == target {
			return i 
		} else if nums[i] > target {
			break
		}
		
	}
	return i
	
}
 // @lc code=end


func main (){
	a := []int{1,3,5,6}
	searchInsert(a, 4)
	fmt.Println("hello")
}


