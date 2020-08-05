/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func reverse(res []int) []int {
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	last := []*TreeNode{root}
	res := make([][]int, 1)
	level := 1
	res[0] = []int{root.Val}
	for len(last) > 0 {
		level++
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
		if level%2 == 0 {
			curRes = reverse(curRes)
		}
		res = append(res, curRes)
	}
	return res[:len(res)-1]
}

// @lc code=end
