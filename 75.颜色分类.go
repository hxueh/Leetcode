/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	left, right, cur := 0, len(nums) - 1, 0
	for cur <= right {
		switch nums[cur] {
		case 2:
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		case 0:
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
            cur++
		default:
			cur++
		}
	}
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
