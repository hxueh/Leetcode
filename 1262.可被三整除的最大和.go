/*
 * @lc app=leetcode.cn id=1262 lang=golang
 *
 * [1262] 可被三整除的最大和
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumDivThree(nums []int) int {
	last := [3]int{0, math.MinInt32, math.MinInt32}
	for _, num := range nums {
		cur := [3]int{0, 0, 0}
		switch num % 3 {
		case 0:
			cur[0] = max(last[0]+num, last[0])
			cur[1] = max(last[1]+num, last[1])
			cur[2] = max(last[2]+num, last[2])
		case 1:
			cur[0] = max(last[2]+num, last[0])
			cur[1] = max(last[0]+num, last[1])
			cur[2] = max(last[1]+num, last[2])
		case 2:
			cur[0] = max(last[1]+num, last[0])
			cur[1] = max(last[2]+num, last[1])
			cur[2] = max(last[0]+num, last[2])
		}
		last = cur
	}
	return last[0]
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
