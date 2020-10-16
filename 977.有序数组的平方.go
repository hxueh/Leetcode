/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sortedSquares(A []int) []int {
	left, right := 0, len(A)-1
	res := make([]int, len(A))
	resIdx := len(A) - 1
	for left <= right {
		var ans int
		if abs(A[left]) < abs(A[right]) {
			ans = A[right] * A[right]
			right--
		} else {
			ans = A[left] * A[left]
			left++
		}
		res[resIdx] = ans
		resIdx--
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
