/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */

// @lc code=start
func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j < i; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
