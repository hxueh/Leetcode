/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{Val: 0}
	cur := &dummyHead
	for l1 != nil || l2 != nil {
		l1Num, l2Num := 9999999999, 9999999999
		if l1 != nil {
			l1Num = l1.Val
		}
		if l2 != nil {
			l2Num = l2.Val
		}
		if l1Num < l2Num {
			cur.Next = &ListNode{Val: l1Num}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2Num}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

// Time complexity O(M+N)
// Space complexity O(1)

// @lc code=end
