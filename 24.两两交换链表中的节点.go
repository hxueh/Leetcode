/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prevNode := dummyHead
	for head != nil && head.Next != nil {
		first, second := head, head.Next
		first.Next, second.Next, prevNode.Next = second.Next, first, second
		prevNode, head = first, first.Next
	}
	return dummyHead.Next
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
