/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
import "sort"

func helper(candidates []int, start int, target int, res *[][]int, curRes []int) {
	if target == 0 {
		*res = append(*res, curRes)
	} else if target > 0 {
		for idx := start; idx < len(candidates); idx++ {
			if idx > start && candidates[idx] == candidates[idx-1] {
				continue
			}
			newCurRes := make([]int, len(curRes)+1)
			copy(newCurRes, curRes)
			newCurRes[len(newCurRes)-1] = candidates[idx]
			helper(candidates, idx+1, target-candidates[idx], res, newCurRes)
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	helper(candidates, 0, target, &res, make([]int, 0))
	return res
}

// @lc code=end
