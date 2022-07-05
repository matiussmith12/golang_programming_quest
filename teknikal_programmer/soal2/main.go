package main

import "fmt"

func displayDeret(T int) {
	x := 1
	for i := 1; i <= T; i++ {
		fmt.Printf("%d ", x * i)
		x *= i
	}
}

func main() {
	displayDeret(6)
}