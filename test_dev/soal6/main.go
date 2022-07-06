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

func countFreqVariant(s string) {
    sMap := make(map[rune]int)
    sOut := ""

    for _, c := range s {
        sMap[c]++
        if sMap[c] == 1 {
            sOut += string(c)
			fmt.Println("SOUT",sOut)
        }
    }

    for _, c := range sOut {
        if sMap[c] > 1 {
            fmt.Print(sMap[c])
        }
        fmt.Printf("%c", c)
    }
}

func main() {
	countFreq("programming")
	countFreqVariant("messi")
}