/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	last := []int{1}
	for i := 1; i <= rowIndex; i++ {
		res := make([]int, i+1)
		res[0], res[len(res)-1] = 1, 1
		for j := 1; j < len(res)-1; j++ {
			res[j] = last[j-1] + last[j]
		}
		last = res
	}
	return last
}

// @lc code=end
