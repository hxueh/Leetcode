/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(matrix [][]int, row int, col int, dp [][]int) {
	ans := 1
	cur := matrix[row][col]
	if row > 0 && matrix[row-1][col] > cur {
		if dp[row-1][col] == 0 {
			helper(matrix, row-1, col, dp)
		}
		ans = max(1+dp[row-1][col], ans)
	}
	if row < len(matrix)-1 && matrix[row+1][col] > cur {
		if dp[row+1][col] == 0 {
			helper(matrix, row+1, col, dp)
		}
		ans = max(1+dp[row+1][col], ans)
	}
	if col > 0 && matrix[row][col-1] > cur {
		if dp[row][col-1] == 0 {
			helper(matrix, row, col-1, dp)
		}
		ans = max(1+dp[row][col-1], ans)
	}
	if col < len(matrix[row])-1 && matrix[row][col+1] > cur {
		if dp[row][col+1] == 0 {
			helper(matrix, row, col+1, dp)
		}
		ans = max(1+dp[row][col+1], ans)
	}
	dp[row][col] = ans
}

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if dp[row][col] == 0 {
				helper(matrix, row, col, dp)
			}
			ans = max(ans, dp[row][col])
		}
	}
	return ans
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
