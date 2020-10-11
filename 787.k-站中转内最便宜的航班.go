/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start

type DstPrice struct {
	dst   int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	dstList := make([][]DstPrice, n)
	for _, flight := range flights {
		dstList[flight[0]] = append(dstList[flight[0]], DstPrice{dst: flight[1], price: flight[2]})
	}
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0
	last := []int{src}
	for i := 0; i <= K; i++ {
		cur := make([]int, 0)
		curPrices := make([]int, n)
		copy(curPrices, prices)
		for _, curSrc := range last {
			curPrice := curPrices[curSrc]
			for _, dstPrice := range dstList[curSrc] {
				if dstPrice.price+curPrice < prices[dstPrice.dst] {
					prices[dstPrice.dst] = dstPrice.price + curPrice
					cur = append(cur, dstPrice.dst)
				}
			}
		}
		last = cur
	}
	if prices[dst] == math.MaxInt32 {
		return -1
	}
	return prices[dst]
}

// Time complexity O(KN)
// Space complexity O(N)

// @lc code=end
