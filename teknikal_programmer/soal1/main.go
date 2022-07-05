package main

import (
	"fmt"
	"strings"
)


func distinctChar(s string){
	lowerString :=strings.ToLower(s)

	freq := make(map[string]int)
	for _, num := range lowerString {
        freq[string(num)]++// = freq[string(num)]+1
	}

	fmt.Println(freq)
}
func main() {
	distinctChar("AdrinaSetiyaiSitrus")
	distinctChar("aaabbbb")
	distinctChar("xyzsaab")

}