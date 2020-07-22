/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	byteMap := make(map[byte]int)
	ans, curAns, start := 0, 0, 0
	for idx := 0; idx < len(s); idx++ {
		curChar := s[idx]
		if beforeIdx, ok := byteMap[curChar]; ok {
			start = max(beforeIdx, start)
			curAns = idx - start
		} else {
			curAns++
		}
		byteMap[curChar] = idx
		ans = max(curAns, ans)
	}
	return ans
}

// Time complexity O(N)
// Space complexity O(Character encoding)

// @lc code=end
