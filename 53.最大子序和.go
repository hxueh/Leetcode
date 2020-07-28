/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur += nums[i]
		res = max(res, cur)
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
