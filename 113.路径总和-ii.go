/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func helper(root *TreeNode, sum int, res *[][]int, curRes []int) {
	if root == nil {
		return
	}
	newCurRes := make([]int, len(curRes)+1)
	copy(newCurRes, curRes)
	newCurRes[len(newCurRes)-1] = root.Val
	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			*res = append(*res, newCurRes)
		}
	} else {
		helper(root.Left, sum-root.Val, res, newCurRes)
		helper(root.Right, sum-root.Val, res, newCurRes)
	}
}
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	helper(root, sum, &res, make([]int, 0))
	return res
}

// @lc code=end
