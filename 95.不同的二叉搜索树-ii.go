/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				res = append(res, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return res
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}

// @lc code=end
