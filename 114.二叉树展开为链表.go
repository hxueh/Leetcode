/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	if root.Left != nil {
		cur := root.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

// @lc code=end
