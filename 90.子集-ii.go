/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
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

	nums = []int{1,2,2}
	fmt.Printf("%v\n", nums)
	fmt.Printf("%v\n", subsetsWithDup(nums))

	// nums = []int{9,0,3,5,7}
	// fmt.Printf("%v\n", nums)
	// fmt.Printf("%v\n", subsets(nums))
}
func subsetsWithDup(nums []int) (ans [][]int) {
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

		if cur > 0 && nums[cur] == nums[cur-1] {
			set = set[:len(set)-2]
			dfs(cur + 1)
		} else {
			set = set[:len(set)-1]
			dfs(cur + 1)
		}
    }
	dfs(0)
	return
}
// @lc code=end

