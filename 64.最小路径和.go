/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
    dp := make([][]int, len(grid))
    for idx := range dp {
        dp[idx] = make([]int, len(grid[idx]))
    }
    dp[0][0] = grid[0][0]
    for i := range dp {
        for j := range dp[i] {
            if i == 0 && j == 0 {
                continue
            } else if i == 0 {
                dp[i][j] = dp[i][j - 1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i - 1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
			}
        }
	}
	return dp[len(dp) - 1][len(dp[0]) - 1]
}
// @lc code=end
