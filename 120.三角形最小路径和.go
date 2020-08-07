/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
package leetcode

func min(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < res {
			res = list[i]
		}
	}
	return res
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] = triangle[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] = min([]int{triangle[i-1][j-1], triangle[i-1][j]}) + triangle[i][j]
		}
		triangle[i][len(triangle[i])-1] = triangle[i-1][len(triangle[i-1])-1] + triangle[i][len(triangle[i])-1]
	}
	return min(triangle[len(triangle)-1])
}

// @lc code=end
