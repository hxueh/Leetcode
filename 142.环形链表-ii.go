/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head.Next
	for fast != nil && fast.Next != nil && fast != slow {
		fast, slow = fast.Next.Next, slow.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

// @lc code=end

