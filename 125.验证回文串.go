/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for (!(unicode.IsLetter(rune(s[left])) || unicode.IsNumber(rune(s[left])))) && left < right {
			left++
		}
		for (!(unicode.IsLetter(rune(s[right])) || unicode.IsNumber(rune(s[right])))) && left < right {
			right--
		}
		if left >= right {
			break
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
