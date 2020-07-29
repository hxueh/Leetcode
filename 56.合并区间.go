/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
import "sort"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	idx := 1
	for idx < len(intervals) {
		first, second := intervals[idx-1], intervals[idx]
		if second[0] <= first[1] {
			intervals[idx-1][1] = max(first[1], second[1])
			intervals = append(intervals[:idx], intervals[idx+1:]...)
		} else {
			idx++
		}
	}
	return intervals
}

// Time complexity O(NlogN)
// Space complexity O(1)

// @lc code=end
