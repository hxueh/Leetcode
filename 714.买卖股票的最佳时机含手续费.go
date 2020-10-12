/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	last := [2]int{-prices[0], 0}
	for i := 1; i < len(prices); i++ {
		cur := [2]int{max(last[0], last[1]-prices[i]), max(last[0]+prices[i]-fee, last[1])}
		last = cur
	}
	return last[1]
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
