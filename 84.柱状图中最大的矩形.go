/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 1)
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	res := 0
	for i := 1; i < len(heights); i++ {
		curHeight, lastHeight := heights[i], heights[stack[len(stack)-1]]
		for len(stack) > 0 && curHeight < lastHeight {
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			res = max(width*lastHeight, res)
			lastHeight = heights[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return res
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
