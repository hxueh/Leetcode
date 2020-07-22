/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start

import "math"

func reverse(x int) int {
	res, less := 0, false
	if x < 0 {
		x, less = -x, true
	}
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	if res > math.MaxInt32 {
		return 0
	}
	if less {
		return -res
	}
	return res
}
// @lc code=end
