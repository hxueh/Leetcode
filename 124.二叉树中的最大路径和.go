/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree root.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := max(helper(root.Left, res), 0)
	right := max(helper(root.Right, res), 0)
	*res = max(*res, root.Val+left+right)
	return root.Val + max(left, right)
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

// @lc code=end
