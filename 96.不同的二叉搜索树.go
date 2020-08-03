/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
package leetcode

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[0], res[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			res[i] += res[j] * res[i-j-1]
		}
	}
	return res[n]
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
