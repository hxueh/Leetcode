/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	last := []*TreeNode{root}
	res := make([][]int, 1)
	res[0] = []int{root.Val}
	for len(last) > 0 {
		curRes := make([]int, 0)
		cur := make([]*TreeNode, 0)
		for _, node := range last {
			if node.Left != nil {
				curRes = append(curRes, node.Left.Val)
				cur = append(cur, node.Left)
			}
			if node.Right != nil {
				curRes = append(curRes, node.Right.Val)
				cur = append(cur, node.Right)
			}
		}
		last = cur
		res = append(res, curRes)
	}
	return res[:len(res)-1]
}

// @lc code=end
