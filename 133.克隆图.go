/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func helper(node *Node, nodeMap map[int]*Node) *Node {
	if newNode, ok := nodeMap[node.Val]; ok {
		return newNode
	}
	newNode := &Node{Val: node.Val}
	nodeMap[node.Val] = newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, helper(neighbor, nodeMap))
	}
	return newNode
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap := make(map[int]*Node)
	return helper(node, nodeMap)
}

// @lc code=end
