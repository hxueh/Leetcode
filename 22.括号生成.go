/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	res := make([]string, 0)
	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis(i) {
			for _, right := range generateParenthesis(n - i - 1) {
				res = append(res, "(" + left + ")" + right)
			}
		}
	}
	return res
}
// @lc code=end
