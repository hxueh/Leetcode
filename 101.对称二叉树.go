/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val == right.Val {
		return helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

// Time complexity O(logN)
// Space complexity O(logN)

// @lc code=end
