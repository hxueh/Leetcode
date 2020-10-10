/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getZeroAndOne(str string) (int, int) {
	var zero, one int
	for _, char := range str {
		switch char {
		case '0':
			zero++
		case '1':
			one++
		}
	}
	return zero, one
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n + 1)
	}
	for _, s := range strs {
		zero, one := getZeroAndOne(s)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i - zero][j - one] + 1)
			}
		}
	}
	return dp[m][n]
}

// Time complexity O(m*n*len(strs))
// Space complexity O(m * n)

// @lc code=end
