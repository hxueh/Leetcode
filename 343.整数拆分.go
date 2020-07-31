/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], dp[i-j]*(i-j))
		}
	}
	return dp[n]
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
