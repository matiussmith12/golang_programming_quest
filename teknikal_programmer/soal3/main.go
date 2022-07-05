package main

import (
	"fmt"
	"math/big"
)


func isPrime(n int64) bool{
	if big.NewInt(n).ProbablyPrime(0) {
		return true
	} else {
		return false
	}
}

func validateValue(T, N int64) bool {
	isValidate := true 

	if N >= 30 {
		isValidate = false
	}
	if !isPrime(N) {
		isValidate = false
	}

	if T < 0 {
		isValidate = false
	}

	return isValidate
}

func nextPrime(n int64) int64{
	if isPrime(n) {
		n += 1
	}

	for !isPrime(n) && n < 30 {
		n += 1
	}

	return n
}

func displayPattern(T, N int64) {
	Ok := validateValue(T, N)

	if Ok {
		var i, j int64

		for i = 0; i < T; i++ {
			temp := N
			for j = 0; j <= i; j++ {
				if(isPrime(temp)){
					fmt.Printf("%d ", temp)
				}
				temp = nextPrime(temp)
			}
			fmt.Println()
		}
	}
}

func main() {
	displayPattern(10, 29)

}