/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func helper(res *[][]int, last []*TreeNode) {
	curLast := make([]*TreeNode, 0)
	curRes := make([]int, 0)
	for _, node := range last {
		if node.Left != nil {
			curRes = append(curRes, node.Left.Val)
			curLast = append(curLast, node.Left)
		}
		if node.Right != nil {
			curRes = append(curRes, node.Right.Val)
			curLast = append(curLast, node.Right)
		}
	}
	if len(curLast) > 0 {
		helper(res, curLast)
		*res = append(*res, curRes)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	helper(&res, []*TreeNode{root})
	res = append(res, []int{root.Val})
	return res
}

// @lc code=end
