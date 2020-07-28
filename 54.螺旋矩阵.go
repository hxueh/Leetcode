/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	for len(matrix) > 0 {
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		for i := 0; i < len(matrix); i++ {
			if len(matrix[i]) > 0 {
				res = append(res, matrix[i][len(matrix[i])-1])
				matrix[i] = matrix[i][:len(matrix[i])-1]
			}
		}
		if len(matrix) == 0 {
			break
		}
		for i := len(matrix[len(matrix)-1]) - 1; i >= 0; i-- {
			res = append(res, matrix[len(matrix)-1][i])
		}
		matrix = matrix[:len(matrix)-1]
		for i := len(matrix) - 1; i >= 0; i-- {
			if len(matrix[i]) > 0 {
				res = append(res, matrix[i][0])
				matrix[i] = matrix[i][1:]
			}
		}
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
