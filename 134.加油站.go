/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	res, remaining, rest := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		remaining += gas[i] - cost[i]
		rest += gas[i] - cost[i]
		if remaining < 0 {
			res = i + 1
			remaining = 0
		}
	}
	if rest < 0 {
		return -1
	}
	return res
}

// @lc code=end
