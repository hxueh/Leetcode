/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			return true
		}
		return false
	}
	sum -= root.Val
	if root.Left == nil {
		return hasPathSum(root.Right, sum)
	}
	if root.Right == nil {
		return hasPathSum(root.Left, sum)
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// @lc code=end
