package main

import "fmt"


func main() {
    var t int
    
    fmt.Scan(&t)
    
    for t > 0{
        var a, b int
		//cnt := 0
        fmt.Scan(&a, &b)
        
        if a % b == 0 {
			fmt.Println(0)
        } else {
			w := a % b
			fmt.Println("W", w)
			fmt.Println("B-W",b - w)
		}
		t--
    }
    
}