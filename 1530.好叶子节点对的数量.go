/*
 * @lc app=leetcode.cn id=1530 lang=golang
 *
 * [1530] 好叶子节点对的数量
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

func helper(root *TreeNode, distance int, res *int) []int {
    if root.Left == nil && root.Right == nil {
        return make([]int, 1)
    }
    left, right := []int{}, []int{}
    if root.Left != nil {
        left = helper(root.Left, distance, res)
    }
    if root.Right != nil {
        right = helper(root.Right, distance, res)
    }
    for i := range left {
        left[i]++
    }
    for i := range right {
        right[i]++
    }
    for _, i := range left {
        for _, j := range right {
            if i + j <= distance {
                *res++
            }
        }
    }
    return append(left, right...)
}

func countPairs(root *TreeNode, distance int) int {
    var res int
    helper(root, distance, &res)
    return res
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
