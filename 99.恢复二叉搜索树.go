/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func getInorder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	left := getInorder(root.Left)
	left = append(left, root)
	right := getInorder(root.Right)
	left = append(left, right...)
	return left
}
func recoverTree(root *TreeNode) {
	inorder := getInorder(root)
	var node1 *TreeNode
	var node2 *TreeNode
	for i := 0; i < len(inorder)-1; i++ {
		if inorder[i].Val > inorder[i+1].Val {
			if node1 == nil {
				node1, node2 = inorder[i], inorder[i+1]
			} else {
				node2 = inorder[i+1]
			}
		}
	}
	node1.Val, node2.Val = node2.Val, node1.Val
}

// @lc code=end
