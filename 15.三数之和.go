/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start

import "sort"

type Three struct {
	first, second, third int
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	resMap := make(map[Three]struct{})
	for idx := 0; idx < len(nums) - 2; idx++ {
		left, right := idx + 1, len(nums) - 1
		for left < right {
			sum := nums[idx] + nums[left] + nums[right]
			if sum == 0 {
				curThree := Three{
					first: nums[idx],
					second: nums[left],
					third: nums[right],
				}
				resMap[curThree] = struct{}{}
                left++
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	for key, _ := range resMap {
		res = append(res, []int{key.first, key.second, key.third})
	}
	return res
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
