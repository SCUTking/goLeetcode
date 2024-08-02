package main

import (
	"fmt"
)

func main() {
	n := 0
	q := 0

	fmt.Scan(&n, &q)
	var count int64
	var sum int64
	for n > 0 {
		n--
		var a int
		fmt.Scan(&a)
		if a == 0 {
			count++
		} else {
			sum += int64(a)
		}
	}

	for q > 0 {
		q--
		var l, r int64
		fmt.Scan(&l, &r)
		fmt.Printf("%d %d\n", sum+count*l, sum+count*r)
	}

}
