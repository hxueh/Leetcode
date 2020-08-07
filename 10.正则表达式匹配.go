/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			return isMatch(s[1:], p) || isMatch(s[1:], p[2:]) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	}
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		return isMatch(s[1:], p[1:])
	}
	return false
}

// @lc code=end
