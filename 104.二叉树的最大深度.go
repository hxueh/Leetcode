/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(helper(root.Left, depth+1), helper(root.Right, depth+1))
}

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}
// @lc code=end
