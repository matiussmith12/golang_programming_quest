package main

import "fmt"

func countFreq(s string)  {
	sMap := make(map[string]int)

	for _, v := range s {
		sMap[string(v)]++
	}

	for i, v := range sMap {
		fmt.Printf("%v%v", i, v)
	}
}
func main() {
	countFreq("programming")
}