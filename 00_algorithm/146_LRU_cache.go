package main

// 146. LRU 缓存机制

// python中有OrderedDict
// java中有LinkedHashMap
type LRUCache struct {
	size int
	capacity int
	cache map[int] *DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkNode(key, value int) *DLinkedNode  {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		capacity: capacity,
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size ++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size --
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}


func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}


func main() {
	
}
