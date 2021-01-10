package main

import (
	"fmt"
)

func main() {
	var nums []int

	nums = []int{0, 1, 2, 4, 5, 7}
	fmt.Printf("%#v\n", summaryRanges(nums))

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Printf("%#v\n", summaryRanges(nums))

	nums = []int{}
	fmt.Printf("%#v\n", summaryRanges(nums))

	nums = []int{-1}
	fmt.Printf("%#v\n", summaryRanges(nums))

	nums = []int{0}
	fmt.Printf("%#v\n", summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res = append(res, []int{nums[i]})
		} else if nums[i] == nums[i-1]+1 {
			lastIndex := len(res) - 1
			res[lastIndex] = append(res[lastIndex], nums[i])
		} else {
			res = append(res, []int{nums[i]})
		}
	}
	var res2 []string
	var tmp string
	for i := 0; i < len(res); i++ {
		if len(res[i]) == 1 {
			tmp = fmt.Sprintf("%d", res[i][0])
			res2 = append(res2, tmp)
		} else {
			tmp = fmt.Sprintf("%d->%d", res[i][0], res[i][len(res[i])-1])
			res2 = append(res2, tmp)
		}
	}
	return res2
}
