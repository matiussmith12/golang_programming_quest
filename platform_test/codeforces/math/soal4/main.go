package main

import (
	"fmt"
	"sort"
)


func main () {
	var a, b, c int
	ans := make([]int, 6)
	fmt.Scan(&a, &b, &c)

	ans[0] = a + b + c
	ans[1] = a * b * c
	ans[2] = a + b * c
	ans[3] = a * b + c 
	ans[4] = (a + b) * c
	ans[5] = a * (b + c)

	sort.Ints(ans)
	fmt.Println(ans[len(ans) - 1])
}