/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
package leetcode

func lengthOfLastWord(s string) int {
	endIdx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && endIdx == -1 {
			endIdx = i
		}
		if s[i] == ' ' && endIdx != -1 {
			return endIdx - i
		}
		if i == 0 && endIdx != -1 {
			return endIdx + 1
		}
	}
	return 0
}

// @lc code=end
