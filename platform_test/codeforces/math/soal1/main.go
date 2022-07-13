package main

import "fmt"


func main () {
	var k, n int64
	fmt.Scan(&n, &k)

	if k <= (n + 1) / 2 {
		fmt.Println(k * 2 - 1)
	} else {
		fmt.Println( 2 * (k - ( n + 1)/ 2))
	}
}