/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
import "sort"

func helper(candidates []int, target int, res *[][]int, curRes []int) {
	if target == 0 {
		*res = append(*res, curRes)
	} else if target > 0 {
		for idx, num := range candidates {
			newCurRes := make([]int, len(curRes)+1)
			copy(newCurRes, curRes)
			newCurRes[len(newCurRes)-1] = num
			helper(candidates[idx:], target-num, res, newCurRes)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	helper(candidates, target, &res, make([]int, 0))
	return res
}

// @lc code=end
