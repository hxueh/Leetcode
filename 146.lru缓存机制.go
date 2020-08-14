/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	dict      map[int]*Node
	len       int
	capacity  int
	dummyHead *Node
	dummyTail *Node
}

// Constructor construct an LRUCache
func Constructor(capacity int) LRUCache {
	dummyHead, dummyTail := &Node{}, &Node{}
	dummyHead.next, dummyTail.prev = dummyTail, dummyHead
	return LRUCache{dict: make(map[int]*Node), capacity: capacity, dummyHead: dummyHead, dummyTail: dummyTail}
}

func (this *LRUCache) moveToTail(node *Node) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) addToTail(node *Node) {
	node.prev, node.next, this.dummyTail.prev.next, this.dummyTail.prev = this.dummyTail.prev, this.dummyTail, node, node
}

func (this *LRUCache) removeHead() {
	node := this.dummyHead.next
	delete(this.dict, node.key)
	this.removeNode(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.dict[key]; !ok {
		return -1
	} else {
		this.moveToTail(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.dict[key]; ok {
		node.val = value
		this.moveToTail(node)
	} else {
		node := &Node{key: key, val: value}
		this.dict[key] = node
		if this.len == this.capacity {
			this.addToTail(node)
			this.removeHead()
		} else {
			this.addToTail(node)
			this.len++
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
