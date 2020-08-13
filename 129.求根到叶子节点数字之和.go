/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func helper(root *TreeNode, curNum int, resList *[]int) {
	curNum = curNum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*resList = append(*resList, curNum)
	} else {
		if root.Left != nil {
			helper(root.Left, curNum, resList)
		}
		if root.Right != nil {
			helper(root.Right, curNum, resList)
		}
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resList := make([]int, 0)
	helper(root, 0, &resList)
	res := 0
	for _, num := range resList {
		res += num
	}
	return res
}

// @lc code=end
