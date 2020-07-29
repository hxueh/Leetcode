/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					obstacleGrid[i][j] = 1
				} else if i == 0 {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				} else if j == 0 {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[len(obstacleGrid)-1])-1]
}

// Time complexity O(M*N)
// Space complexity O(M*N)

// @lc code=end
