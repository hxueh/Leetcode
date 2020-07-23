/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		leftNum, rightNum, midNum := nums[left], nums[right], nums[mid]
		if midNum == target {
			return mid
		}
		if midNum > rightNum {
			if target > midNum {
				left = mid + 1
			} else {
				if target >= leftNum {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else {
			if target < midNum {
				right = mid - 1
			} else {
				if target <= rightNum {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return -1
}

// Time complexity O(log(n))
// Space complexity O(1)

// @lc code=end
