/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package main

import "fmt"

func main() {
	var nums []int

	// nums = []int{}
	// fmt.Printf("%v, %v\n", nums, subsets(nums))


	// nums = []int{1,2,3}
	// fmt.Printf("%v, %v\n", nums, subsets(nums))

	nums = []int{1,2,3,4}
	fmt.Printf("%v\n", nums)
	fmt.Printf("%v\n", subsets(nums))

	// nums = []int{9,0,3,5,7}
	// fmt.Printf("%v\n", nums)
	// fmt.Printf("%v\n", subsets(nums))
}

// a := []int


func subsets(nums []int) (ans [][]int) {

	set := []int{}
	fmt.Printf("set: %v\n", set)
    var dfs func(int)
    dfs = func(cur int) {
        if cur == len(nums) {
			fmt.Printf("cur: %d, set: %v\n", cur, set)
			tmp := append([]int(nil), set...) // equally, copy
            ans = append(ans, tmp)
            return
        }
        set = append(set, nums[cur])
        dfs(cur + 1)
        set = set[:len(set)-1]
        dfs(cur + 1)
    }
	dfs(0)
	
    return


	res := [][]int{[]int{}}

	for i := 0 ; i < len(nums) ; i++ {
		// tmp := make([][]int, len(res))
		// fmt.Printf("i: %d, res: size %d, %v\n", i, len(res), res)
		// copy(tmp, res)
		next := [][]int{}
		for j := 0 ; j < len(res) ; j++ {
			tmp := make([]int, len(res[j]))
			copy(tmp, res[j])
			// fmt.Printf("j: %d, nums[i]: %d, , tmp[j]: %v", j, nums[i], tmp[j])
			tmp = append(tmp, nums[i])
			// fmt.Printf(", tmp[j]: %v\n", tmp[j])
			next = append(next, tmp)
		}
		res = append(res, next...)
	}
	return res
}


// @lc code=end

