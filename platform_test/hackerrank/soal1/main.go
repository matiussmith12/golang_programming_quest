package main

import (
	"fmt"
	"sort"
)

func countStudents(height []int) int {
	count := 0
	newOrdered := []int  {}

	newOrdered = append(newOrdered, height...)
	sort.Ints(newOrdered)

	for i := 0 ; i < len(height); i++ {
		if newOrdered[i] != height[i] {
			count++
		}
	}

	return count
}

func main() {
	unOrderedHeight := []int {1, 1, 3, 4, 1}
	fmt.Println(countStudents(unOrderedHeight))
}