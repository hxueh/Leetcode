/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	half := myPow(x, n / 2)
	return half * half * myPow(x, n % 2)
}

// Time complexity O(logN)
// Space complexity O(logN)

// @lc code=end
