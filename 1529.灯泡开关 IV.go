/*
 * @lc app=leetcode.cn id=1529 lang=golang
 *
 * [1529] 灯泡开关 IV
 */

// @lc code=start
func minFlips(target string) int {
	last, res := '0', 0
	for _, char := range target {
		if char != last {
			last, res = char, res + 1
		}
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
