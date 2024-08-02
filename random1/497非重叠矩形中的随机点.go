package random1

import (
	"math/rand"
	"sort"
)

//type Solution struct {
//	inNodes [][]int
//}
//
//func Constructor(rects [][]int) Solution {
//	solution := Solution{inNodes: make([][]int, 0)}
//	for _, rect := range rects {
//		leftX := rect[0]
//		leftY := rect[1]
//		rightX := rect[2]
//		rightY := rect[3]
//
//		for i := leftX; i < rightX; i++ {
//			for j := leftY; j < rightY; j++ {
//				temp := []int{i, j}
//				solution.inNodes = append(solution.inNodes, temp)
//			}
//		}
//	}
//
//	return solution
//}
//
//func (this *Solution) Pick() []int {
//	intn := rand.Intn(len(this.inNodes))
//	return this.inNodes[intn]
//}

type Solution struct {
	rects [][]int
	sum   []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)

	}
	return Solution{rects: rects, sum: sum}

}

func (this *Solution) Pick() []int {
	k := rand.Intn(this.sum[len(this.sum)-1])
	//k+1是防止变为0
	rectIndex := sort.SearchInts(this.sum, k+1) - 1
	r := this.rects[rectIndex]
	a, b, y := r[0], r[1], r[2]
	da := (k - this.sum[rectIndex]) / (y - b + 1)
	db := (k - this.sum[rectIndex]) % (y - b + 1)
	return []int{a + da, b + db}

}
