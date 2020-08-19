/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		odd, even := helper(s, i, i), helper(s, i, i+1)
		res += (odd + even)
	}
	return res
}

func helper(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	length := right - left - 1
	return length/2 + length%2
}

// @lc code=end
