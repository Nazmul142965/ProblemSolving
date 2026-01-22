package main

import "fmt"

func main() {
    var t int

    fmt.Scanf("%d", &t)

    // Loop for each test case
    for t > 0 {
        var x, y int
        fmt.Scanf("%d %d", &x, &y)

        sum := x+y
		
        if (sum > 6){
            fmt.Println("yes")
        }else{
            fmt.Println("no")
        }

        t--
    }
}
