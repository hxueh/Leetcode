/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][len(res[i])-1] = 1, 1
		if i > 0 {
			for j := 1; j < len(res[i])-1; j++ {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}

// @lc code=end
