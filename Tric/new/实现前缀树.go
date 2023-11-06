package new

type Trie struct {
	isEnd int
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, i := range word {
		i2 := i - 'a'
		if root.next[i2] == nil {
			root.next[i2] = &Trie{}
		}
		root = root.next[i2]
	}
	root.isEnd++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for _, i := range word {
		i2 := i - 'a'
		if root.next[i2] == nil {
			return false
		}
		root = root.next[i2]
	}
	if root.isEnd == 0 {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, i := range prefix {
		i2 := i - 'a'
		if root.next[i2] == nil {
			return false
		}
		root = root.next[i2]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
