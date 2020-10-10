/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
