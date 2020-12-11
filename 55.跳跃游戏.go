/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
    n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
			if i <= rightmost {
					rightmost = max(rightmost, i + nums[i])
					if rightmost >= n - 1 {
							return true
					}
			}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

