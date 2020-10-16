/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		middle := (left + right) / 2
		if nums[middle] < nums[right] {
			right = middle
		} else if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right--
		}
	}
	return nums[left]
}

// Time complexity O(logN)
// Space complexity O(1)

// @lc code=end
