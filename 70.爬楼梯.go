/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	lastZero, lastOne := 1, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = lastOne + lastZero
		lastZero, lastOne = lastOne, res
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
