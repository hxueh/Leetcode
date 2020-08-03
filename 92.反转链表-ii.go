/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummyHead := &ListNode{Next: head}
	mHead, beforeMHead := dummyHead, dummyHead
	for i := 0; i < m; i++ {
		beforeMHead = mHead
		mHead = mHead.Next
	}
	dummyBeforeMHead := beforeMHead
	nTail, cur := mHead, mHead
	for i := 0; i <= n-m; i++ {
		nTail = cur
		cur.Next, dummyBeforeMHead, cur = dummyBeforeMHead, cur, cur.Next
	}
	mHead.Next = cur
	beforeMHead.Next = nTail
	return dummyHead.Next
}

// Time complexity O(M+N)
// Space complexity O(1)

// @lc code=end
