/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	inorderIdx := 0
	for inorderIdx < len(inorder) && inorder[inorderIdx] != head.Val {
		inorderIdx++
	}
	head.Left = buildTree(preorder[:inorderIdx], inorder[:inorderIdx])
	head.Right = buildTree(preorder[inorderIdx:], inorder[inorderIdx+1:])
	return head
}

// @lc code=end
