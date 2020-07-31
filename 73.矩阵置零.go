/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
    firstIsZero := false
    if matrix[0][0] == 0 {
        firstIsZero = true
    }
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstIsZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
// @lc code=end
