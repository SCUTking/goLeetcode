package training_camp

// 不能使用这种方法  因为index加入之后就不会发生改变了
type entry struct {
	index int
	value int
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}
type LinkedNode struct {
	pre, next  *LinkedNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*LinkedNode, 0),
		capacity: capacity,
		head:     NewNode(0, 0),
		tail:     NewNode(0, 0),
	}

	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}

func NewNode(key, value int) *LinkedNode {
	return &LinkedNode{key: key, value: value}
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	next := this.head.next
	next.pre = node
	node.next = next

	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) removeNode(node *LinkedNode) {

	pre := node.pre
	pre.next = node.next
	node.next.pre = pre

}

func (this *LRUCache) MoveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LinkedNode {

	node := this.tail.pre

	this.removeNode(node)

	return node
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.cache[key]

	if !ok {
		return -1
	}
	//调整LRU
	this.MoveToHead(res)
	return res.value
}

// Put 有三种逻辑
// key存在就进行更新
// 不存在，就添加
// 如果添加超过容量，得进行淘汰
func (this *LRUCache) Put(key int, value int) {
	res, ok := this.cache[key]
	if ok {
		this.cache[key].value = value
		this.MoveToHead(res)
		return
	}

	//再添加就超过了
	if this.capacity == len(this.cache) {
		tail := this.removeTail()
		delete(this.cache, tail.key)
	}
	//淘不淘汰都得添加
	node := NewNode(key, value)
	this.cache[key] = node
	this.addToHead(node)
}
