package main

import "fmt"

func countFreq(s string)  {
	sMap := make(map[rune]int)

	for _, v := range s {
		sMap[(v)]++
	}

	for _, v := range s {
		count := sMap[v]
		if count == 0 {
			continue // Char already printed and removed
		}
		delete(sMap, v)
		if count > 1 {
			fmt.Print(count)
		}
		fmt.Print(string(v))
	}
}
func main() {
	countFreq("programming")
}