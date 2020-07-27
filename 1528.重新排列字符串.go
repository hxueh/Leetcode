/*
 * @lc app=leetcode.cn id=1528 lang=golang
 *
 * [1528] 重新排列字符串
 */

// @lc code=start
func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, n := 0, len(s); i < n; i++ {
		res[indices[i]] = s[i]
	}
	return string(res)
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
