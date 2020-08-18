/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var beforeSlow *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		beforeSlow, slow, fast = slow, slow.Next, fast.Next.Next
	}
	beforeSlow.Next = nil
	return &TreeNode{Val: slow.Val, Left: sortedListToBST(head), Right: sortedListToBST(slow.Next)}
}

// @lc code=end
