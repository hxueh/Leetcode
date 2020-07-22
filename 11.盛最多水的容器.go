/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	res := 0
	for left < right {
		curRes := min(height[left], height[right]) * (right - left)
		res = max(curRes, res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
