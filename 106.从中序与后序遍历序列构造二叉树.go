/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	head := &TreeNode{Val: postorder[len(postorder)-1]}
	postorder = postorder[:len(postorder)-1]
	idx := 0
	for idx < len(inorder) && inorder[idx] != head.Val {
		idx++
	}
	head.Left = buildTree(inorder[:idx], postorder[:idx])
	head.Right = buildTree(inorder[idx+1:], postorder[idx:])
	return head
}

// @lc code=end
