package training_camp

type LRUCache struct {
	cache  []int       //缓存队列
	m      map[int]int //存储实际数据的，
	maxLen int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make([]int, 0), m: make(map[int]int, capacity), maxLen: capacity}
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.m[key]

	if !ok {
		return -1
	}

	//调整LRU
	for i := 0; i < len(this.cache); i++ {
		if this.cache[i] == key {
			this.cache = append(this.cache[:i], this.cache[i+1:]...)
			this.cache = append(this.cache, key)
			break
		}
	}

	return res
}

// Put 有三种逻辑
// key存在就进行更新
// 不存在，就添加
// 如果添加超过容量，得进行淘汰
func (this *LRUCache) Put(key int, value int) {
	_, ok := this.m[key]
	if ok {
		this.m[key] = value
		this.Get(key) //更新也算访问
		return
	}

	//再添加就超过了
	if this.maxLen == len(this.m) {
		delete(this.m, this.cache[0])
		this.cache = this.cache[1:]
	}
	//淘不淘汰都得更新
	this.m[key] = value
	this.cache = append(this.cache, key)
}
