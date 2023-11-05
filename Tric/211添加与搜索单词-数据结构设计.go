package Tric

type WordDictionary struct {
	next  []*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
	dictionaries := make([]*WordDictionary, 26)
	return WordDictionary{next: dictionaries}
}

func (this *WordDictionary) AddWord(word string) {

	nowNode := this
	for _, v := range word {
		i := v - 'a'
		if nowNode.next[i] == nil {
			dictionaries := make([]*WordDictionary, 26)
			w := &WordDictionary{next: dictionaries}
			nowNode.next[i] = w
		}
		nowNode = nowNode.next[i]
	}
	nowNode.isEnd = true

}

func (this *WordDictionary) Search(word string) bool {

	nowNode := this
	for j, v := range word {
		if v == '.' {
			runes := []rune(word)
			if len(runes) == 1 {
				if nowNode.isEnd {
					return true
				} else {
					return false
				}

			}
			word = string(runes[j+1:])
			if len(nowNode.next) > 0 {
				res := false
				for _, each1 := range nowNode.next {
					for _, each := range each1.next {
						if each != nil {
							search := each.Search(word)
							if search {
								res = true
							}
						}

					}
				}

				return res
			}
		}

		i := v - 'a'

		if nowNode.next[i] == nil {
			return false
		}
		nowNode = nowNode.next[i]
	}
	if nowNode.isEnd {
		return true
	}
	return false

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
