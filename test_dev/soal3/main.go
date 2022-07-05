package main

import (
	"fmt"
	"math/big"
)

func isPrime(n int64) bool {
	if big.NewInt(n).ProbablyPrime(0) {
		return true
	} else {
		return false
	}
}

func main() {
	var i int64
	count := 0
	for i = 1; i <= 100; i++ {
		
		if isPrime(i) {
			fmt.Println("Bilangan Prima: ", i)
		} else if i % 9 == 0 {
			count++
			fmt.Println("Kelipatan 9 ke - ", count)
		} else {
			fmt.Println(i)
		}
	}
}