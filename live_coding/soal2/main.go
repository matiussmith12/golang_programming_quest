package main

import "fmt"

func seive(number int) int {
	// Create Array of boolean with false value
	primeArray := make([]bool, number+1)

	for i := 0; i < number+1; i++ {
		primeArray[i] = false
	}

	answer := 0
	for p := 2; p*p <= number; p++ {
		if !primeArray[p] {
			for j := p*p; j <= number; j+=p {
				if j == p*p {
					answer++
				}
				primeArray[j] = true
			}
		}
	}

	return answer
}



func main() {
	var a int
	fmt.Scan(&a)

	result := seive(a)
	fmt.Println(result)
}