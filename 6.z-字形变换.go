/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 1 || len(s) < 1 {
		return s
	}
	resSlice := make([][]byte, numRows)
	for i := range resSlice {
		resSlice[i] = make([]byte, 0)
	}
	for i := 0; i < len(s); i++ {
		line := 0
		if numRows > 1 {
			line = i % (2*numRows - 2)
		}
		if line > numRows-1 {
			line = (2*numRows - 2) - line
		}
		resSlice[line] = append(resSlice[line], s[i])
	}
	res := ""
	for _, b := range resSlice {
		res += string(b)
	}
	return res
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
