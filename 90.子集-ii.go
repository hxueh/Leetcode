/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 1)
	lastStart := 0
	for i, num := range nums {
		resLen := len(res)
		start := 0
		if i > 0 && num == nums[i-1] {
			start = lastStart
		}
		for j := start; j < resLen; j++ {
			curRes := make([]int, len(res[j])+1)
			copy(curRes, res[j])
			curRes[len(curRes)-1] = num
			res = append(res, curRes)
		}
		lastStart = resLen
	}
	return res
}

// @lc code=end
