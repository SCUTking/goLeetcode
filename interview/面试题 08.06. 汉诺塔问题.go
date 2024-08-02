package interview

// 递归
// 将移动n个盘子转化为移动n-1个盘子，与1个盘子
// 从小问题做起的，1与2这些，然后看看是否能找到规律
func hanota(A []int, B []int, C []int) []int {
	var helper func(n int, a, b, c *[]int) //n是移动元素个数，a是要移动栈，b是辅助，c移动的目标栈
	helper = func(n int, a, b, c *[]int) {
		if n == 1 {
			*c = append(*c, (*a)[len(*a)-1])
			*a = (*a)[:len(*a)-1]
			return
		}
		helper(n-1, a, c, b) //将底层的n-1通过c 移动到b
		//下面两行将a最后一层移动到c
		*c = append(*c, (*a)[len(*a)-1])
		*a = (*a)[:len(*a)-1]
		//因为之前n-1行移动到了b，所以起点为b，辅助为a，目标还是c
		helper(n-1, b, a, c)
	}
	helper(len(A), &A, &B, &C)
	return C
}
