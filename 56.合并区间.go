/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

type sortSlice [][]int

func (a sortSlice) Len() int {
	return len(a)
}

func (a sortSlice) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	} else if a[i][0] < a[j][0] {
		return true
	} else {
		return false
	}
}

func (a sortSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}

	sort.Sort(sortSlice(intervals))
	fmt.Printf("sort is %v\n", intervals)
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] == last[0] {
			last[1] = intervals[i][1]
			continue
		}
		if intervals[i][0] <= last[1] && intervals[i][1] >= last[1] {
			last[1] = intervals[i][1]
			continue
		}

		res = append(res, intervals[i])
	}
	return res
}

func main() {
	var a [][]int
	var b [][]int

	a = [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	fmt.Printf("%v\n", a)
	b = merge(a)
	fmt.Printf("%v\n", b)
	fmt.Println()

	a = [][]int{{2, 6}}
	fmt.Printf("%v\n", a)
	b = merge(a)
	fmt.Printf("%v\n", b)
	fmt.Println()

	a = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {1, 100}}
	fmt.Printf("%v\n", a)
	b = merge(a)
	fmt.Printf("%v\n", b)
	fmt.Println()

	a = [][]int{{1, 3}, {2, 2}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("%v\n", a)
	b = merge(a)
	fmt.Printf("%v\n", b)
	fmt.Println()

	a = [][]int{}
	fmt.Printf("%v\n", a)
	b = merge(a)
	fmt.Printf("%v\n", b)
	fmt.Println()
}

// @lc code=end
