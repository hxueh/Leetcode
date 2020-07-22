/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, dummyHead := 0, ListNode{Val: 0}
	cur := &dummyHead
	for l1 != nil || l2 != nil || carry != 0 {
		l1Num, l2Num := 0, 0
		if l1 != nil {
			l1Num = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Num = l2.Val
			l2 = l2.Next
		}
		curNum := l1Num + l2Num + carry
		carry = curNum / 10
		curNode := ListNode{Val: curNum % 10}
		cur.Next = &curNode
		cur = cur.Next
	}
	return dummyHead.Next
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
