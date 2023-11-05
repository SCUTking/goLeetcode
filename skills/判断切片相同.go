package skills

//go 语言中可以使用反射 reflect.DeepEqual(a, b) 判断 a、b 两个切片是否相等，
//但是通常不推荐这么做，使用反射非常影响性能。
//通常采用的方式如下，遍历比较切片中的每一个元素（注意处理越界的情况）。

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	//有一种情况，两个都是nil的时候，应该是返回true的
	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
