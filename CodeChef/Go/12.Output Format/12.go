package main

import "fmt"

func main() {
    var t int

    fmt.Scanf("%d", &t)

    // Loop for each test case
    for t > 0 {
        var a, b int
        fmt.Scanf("%d %d", &a, &b)
        
        c:= a+b
        fmt.Println(c)
        
        t--
    }
    
}