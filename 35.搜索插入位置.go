/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums) - 1] < target {
		return len(nums)
	}
	left, right := 0, len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target && nums[mid + 1] >= target {
			return mid + 1
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

// Time complexity O(log(n))
// Space complexity O(1)

// @lc code=end
