/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start

func newAppend(curRes []int, num int) []int {
	newRes := make([]int, len(curRes)+1)
	copy(newRes, curRes)
	newRes[len(curRes)] = num
	return newRes
}

func helper(nums []int, start int, res *[][]int, curRes []int) {
	visited := make(map[int]struct{})
	for i := start; i < len(nums); i++ {
		curNum := nums[i]
		if _, ok := visited[curNum]; ok {
			continue
		}
		visited[curNum] = struct{}{}
		if len(curRes) == 0 {
			helper(nums, i+1, res, newAppend(curRes, curNum))
		} else {
			if curNum >= curRes[len(curRes)-1] {
				newRes := newAppend(curRes, curNum)
				*res = append(*res, newRes)
				helper(nums, i+1, res, newRes)
			}
		}
	}
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	helper(nums, 0, &res, make([]int, 0))
	return res
}

// @lc code=end
