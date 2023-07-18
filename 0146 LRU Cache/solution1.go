// Your runtime beats 82.1 % of golang submissions (578 ms)
// Your memory usage beats 41.19 % of golang submissions (79.5 MB)
//
// Hash 表 + 双向链表
// 使用 Hash 表维护缓存
// 使用双向链表维护最近使用顺序，头节点为最近访问的，尾节点为最久使用的
// 当读取记录时，将节点移动到头节点
// 当插入新记录时，将新的记录放到链表的头节点，并记录到 Hash 表中
// 当达到容量上限时再加入新值时需要将最久未使用的记录删除，即将链表尾节点删除，Hash 表删除对应的记录

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int           // 容量
	size     int           // 目前容纳的缓存数量
	cache    map[int]*Node // 用于记录缓存的 Hash 表
	head     *Node         // 头节点
	tail     *Node         // 尾节点
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// RemoveNode 将节点从链表中删除
func (this *LRUCache) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

// RemoveTail 将最后一个节点从链表中删除
func (this *LRUCache) RemoveTail() *Node {
	if this.size > 0 {
		tailNode := this.tail.prev
		this.RemoveNode(tailNode)
		return tailNode
	}
	return nil
}

// InsertHead 将一个节点插入头部
func (this *LRUCache) InsertHead(node *Node) {
	nextNode := this.head.next
	nextNode.prev = node
	this.head.next = node
	node.prev = this.head
	node.next = nextNode
}

// Get 根据 key 获取缓存
func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// 更新访问顺序，将读取的节点移到头部
	this.RemoveNode(node)
	this.InsertHead(node)
	return node.val
}

// Put 记录缓存信息
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		// 当 key 已经存在时
		// 更新缓存
		node.val = value
		// 将节点移动到链表头部
		this.RemoveNode(node)
		this.InsertHead(node)
	} else {
		// 当 key 不存在时
		if this.size == this.capacity {
			// 当容量达到上限
			// 需要删除的缓存为链表的最后一个节点
			tailNode := this.RemoveTail()
			// Hash 表中删除缓存
			delete(this.cache, tailNode.key)
			this.size -= 1
		}

		// 创建缓存记录
		node := &Node{key: key, val: value}
		// 插入到链表头部
		this.InsertHead(node)
		// 记录到 Hash 表中
		this.cache[key] = node
		this.size += 1
	}
}
