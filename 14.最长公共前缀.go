/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for idx := 0; idx < len(strs[0]); idx++ {
		curChar := strs[0][idx]
		for _, str := range strs[1:] {
			if idx >= len(str) {
				return str[:idx]
			}
			if str[idx] != curChar {
				return str[:idx]
			}
		}
	}
	return strs[0]
}
// @lc code=end
