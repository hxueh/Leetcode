/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for idx, num := range nums {
		wanted := target - num
		if wantedIdx, ok := numMap[wanted]; ok {
			return []int{wantedIdx, idx}
		}
		numMap[num] = idx
	}
	return []int{}
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
