/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	cur := 0
	for i, num := range nums {
		cur = max(cur-1, num)
		if cur <= 0 && i < len(nums)-1 {
			return false
		}
	}
	return true
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
