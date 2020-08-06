/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, -1
	} else {
		leftIsBalanced, leftHeight := helper(root.Left)
		if !leftIsBalanced {
			return false, 0
		}
		rightIsBalanced, rightHeight := helper(root.Right)
		if !rightIsBalanced {
			return false, 0
		}
		return abs(leftHeight-rightHeight) <= 1, 1 + max(leftHeight, rightHeight)
	}
}

func isBalanced(root *TreeNode) bool {
	res, _ := helper(root)
	return res
}

// @lc code=end
