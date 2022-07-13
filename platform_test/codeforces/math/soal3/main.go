package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	ans := make([]int64, n)
	var sum float64 = 0

	var i int64
	for i = 0; i < n; i++ {
		fmt.Scan(&ans[i])
		sum = sum + float64(ans[i])
	}

	fmt.Printf("%.12f", sum / float64(n))
}