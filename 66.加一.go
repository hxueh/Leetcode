/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
		curRes := digits[i] + carry
		digits[i], carry = curRes%10, curRes/10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
