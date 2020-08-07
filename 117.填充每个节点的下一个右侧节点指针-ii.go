/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curNext := root.Next
	if root.Left != nil {
		if root.Right == nil {
			for curNext != nil && (curNext.Left == nil && curNext.Right == nil) {
				curNext = curNext.Next
			}
			if curNext != nil {
				if curNext.Left != nil {
					root.Left.Next = curNext.Left
				} else {
					root.Left.Next = curNext.Right
				}
			}
		} else {
			root.Left.Next = root.Right
		}
	}
	if root.Right != nil {
		for curNext != nil && (curNext.Left == nil && curNext.Right == nil) {
			curNext = curNext.Next
		}
		if curNext != nil {
			if curNext.Left != nil {
				root.Right.Next = curNext.Left
			} else {
				root.Right.Next = curNext.Right
			}
		}
	}
	connect(root.Right)
	connect(root.Left)
	return root
}

// @lc code=end
