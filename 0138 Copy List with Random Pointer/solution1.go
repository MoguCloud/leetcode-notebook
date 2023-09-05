// Your runtime beats 63.66 % of golang submissions (2 ms)
// Your memory usage beats 28.38 % of golang submissions (3.52 MB)
//
// Hash Table
// 遍历链表，根据当前节点创建一个值一样的节点，然后存入哈希表，key为原节点指针，value为新节点指针
// 再次遍历链表，同时创建新链表，根据原节点指针找到对应的新节点和random指针对应的新节点
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func copyRandomList(head *Node) *Node {
	memo := make(map[*Node]*Node)
	for node := head; node != nil; node = node.Next {
		if _, ok := memo[node]; ok {
			continue
		}
		newNode := &Node{Val: node.Val}
		memo[node] = newNode
	}
	dummy := &Node{}
	newHead := dummy
	for node := head; node != nil; node = node.Next {
		newNode := memo[node]
		if node.Random != nil {
			newRandom := memo[node.Random]
			newNode.Random = newRandom
		}
		newHead.Next = newNode
		newHead = newHead.Next
	}
	return dummy.Next
}
