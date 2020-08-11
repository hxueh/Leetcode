/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	res := 0
	minPrice := 999999999
	for _, price := range prices {
		minPrice = min(price, minPrice)
		res = max(price-minPrice, res)
	}
	return res
}

// @lc code=end
