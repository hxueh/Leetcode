/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix[0]), len(matrix)
	left, right := 0, m * n - 1
	for left <= right {
		mid := (left + right) / 2
		midI, midJ := mid / m, mid % m
		if matrix[midI][midJ] == target {
			return true
		} else if matrix[midI][midJ] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// Time complexity O(logN)
// Space complexity O(1)

// @lc code=end
