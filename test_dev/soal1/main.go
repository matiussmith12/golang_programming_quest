package main

import "fmt"

func swapValueWithoutTemp(A, B int) (rA, rB int){
	rA = A + B - A
	rB = B + A - B

	return rA, rB
}
func main() {
	fmt.Println(swapValueWithoutTemp(30, 50))
}
