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
    res := 1.0
    for n > 0 {
        if n % 2 > 0 {
            res = x * res
        }
        x *= x
        n /= 2
    }
	return res
}

// Time complexity O(logN)
// Space complexity O(1)

// @lc code=end
