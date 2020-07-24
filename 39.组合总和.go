/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

import (
	"reflect"
	"sort"
)

func helper(candidates []int, target int, res *[][]int, curRes []int) {
	if target == 0 {
		for _, i := range *res {
			if reflect.DeepEqual(i, curRes) {
				return
			}
		}
		*res = append(*res, curRes)
	} else if target > 0 {
		for idx, i := range candidates {
			newCurRes := append(curRes, i)
			helper(candidates[idx:], target-i, res, newCurRes)
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
