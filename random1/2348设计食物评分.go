package random1

import "container/heap"

type pair struct {
	r int
	s string
}

type FoodRatings struct {
	fs map[string]pair
	cs map[string]*hp
}

func constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fs := map[string]pair{}
	cs := map[string]*hp{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fs[f] = pair{r, c}
		if cs[c] == nil {
			cs[c] = &hp{}
		}
		heap.Push(cs[c], pair{r, f})
	}
	return FoodRatings{fs, cs}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.fs[food]
	heap.Push(r.cs[p.s], pair{newRating, food}) // 直接添加新数据，后面 highestRated 再删除旧的
	p.r = newRating
	r.fs[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	h := *r.cs[cuisine]
	for h.Len() > 0 && h[0].r != r.fs[h[0].s].r { // 堆顶的食物评分不等于其实际值
		heap.Pop(&h)
	}
	return h[0].s
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.r > b.r || a.r == b.r && a.s < b.s }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
