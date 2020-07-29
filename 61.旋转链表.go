/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	lenOfListNode, tail := 1, head
	for tail.Next != nil {
		lenOfListNode++
		tail = tail.Next
	}
	k = lenOfListNode - k % lenOfListNode
	tail.Next = head
	for i := 0; i < k; i++ {
		tail = head
		head = head.Next
	}
	tail.Next = nil
	return head
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
