/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	curRes := []int{root.Val}
	res := inorderTraversal(root.Left)
	res = append(res, curRes...)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}
// @lc code=end
