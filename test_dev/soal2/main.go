package main

import "fmt"


func reverse(s string) (result string ){
	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func reverseWord(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}

func main() {
	fmt.Println(reverse("adrian"))
	fmt.Println(reverseWord("messi"))
}