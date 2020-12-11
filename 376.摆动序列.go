/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func max(nums ...int) int {
	var res int = nums[0]
	for _, num := range nums[1:] {
		if num > res {
			res = num
		}
	}
	return res
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dp[i] = 1
		} else if nums[i] > nums[i-1] {
			dp[i] = 2
			for j := i - 1; j >= 1; j-- {
				if nums[j] < nums[j-1] {
					dp[i] = dp[j] + 1
					break
				}
			}
		} else if nums[i] < nums[i-1] {
			dp[i] = 2
			for j := i - 1; j >= 1; j-- {
				if nums[j] > nums[j-1] {
					dp[i] = dp[j] + 1
					break
				}
			}
		}
	}
	return max(dp...)
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
