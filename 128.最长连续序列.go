/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func helper(num int, numsMap map[int]int) {
	if numsMap[num] == 1 {
		helper(num-1, numsMap)
		numsMap[num] += numsMap[num-1]
	}
}

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num] = 1
	}
	for num := range numsMap {
		helper(num, numsMap)
	}
	res := 0
	for _, value := range numsMap {
		if value > res {
			res = value
		}
	}
	return res
}

// @lc code=end
