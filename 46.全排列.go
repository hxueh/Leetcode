/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func copyAndPop(nums []int, popIdx int) []int {
	res := make([]int, len(nums) - 1)
	copy(res, nums[:popIdx])
	copy(res[popIdx:], nums[popIdx + 1:])
	return res
}

func helper(nums []int, curRes []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, curRes)
	} else {
		for idx, num := range nums {
			newNums := copyAndPop(nums, idx)
			newCurRes := append(curRes, num)
			helper(newNums, newCurRes, res)
		}
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
    curRes := make([]int, 0)
	helper(nums, curRes, &res)
	return res
}

// Time complexity O(Nn)
// Space complexity O(Nn)

// @lc code=end
