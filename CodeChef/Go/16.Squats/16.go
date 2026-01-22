package main

import "fmt"

func main() {
    var t int

    fmt.Scanf("%d", &t)

    // Loop for each test case
    for t > 0 {
        var x int
        fmt.Scanf("%d", &x)
        c:= x*15
        fmt.Println(c)

        t--
    }
}
