/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start

package leetcode

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	// stack is min stack
	res, stack := 0, make([]int, 1)
	for idx, h := range height {
		for len(stack) > 1 && height[stack[len(stack)-1]] < h {
			curHeight := min(h, height[stack[0]]) - height[stack[len(stack)-1]]
			curWidth := stack[len(stack)-1] - stack[len(stack)-2]
			curArea := curHeight * curWidth
			res += curArea
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 1 && height[stack[0]] <= h {
			stack = append(stack[1:], idx)
		} else {
			stack = append(stack, idx)
		}
	}
	return res
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
