/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
package main

import "sort"
import "fmt"

func main() {
	var a, b []int	
	// a = []int{2,3,1,3,2,4,6,7,9,2,19}
	// b = []int{2,1,4,3,9,6}
	a = []int{28,6,22,8,44,17}
    b = []int{22,28,8,6}
	c := relativeSortArray(a,b)	
	fmt.Println(c)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res1 := []int{}
	res2 := []int{}
	fmt.Printf("%v, %v\n", arr1, arr2)
	copyarr2 := append([]int{}, arr2...)

	sort.Ints(copyarr2)
	fmt.Printf("copyarr2, %v\n", copyarr2)
	map1 := map[int]int{}

	for i := 0 ; i < len(arr1) ; i++ {
		if	ok := find(copyarr2, arr1[i]) ; ok {
			map1[arr1[i]]++
		} else {
			res2 = append(res2, arr1[i])
		}
	}
	fmt.Printf("res2 is %v\n",res2)
	for i := 0 ; i < len(arr2) ; i ++ {
		size, _ := map1[arr2[i]]
		for j := 0 ; j < size ; j++ {
			res1 = append(res1, arr2[i])
		}
	}
	sort.Ints(res2)
	return append(res1, res2...)
}

func find(a []int, target int) bool {
	// fmt.Println("hello")
	if len(a) < 1 {
		return false
	}
	var mid int
	down := 0 
	upper := len(a)	 - 1
	for down <= upper {
		// mid = (down + upper ) >> 2
		mid = down + (upper - down) >> 1
		if a[mid] == target {
			return true
		}	
		if target > a[mid] {
			down = mid + 1
		} else {
			upper = mid - 1
		}
	}
	return false
}
// @lc code=end

