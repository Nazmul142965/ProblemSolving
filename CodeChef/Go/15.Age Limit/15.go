package main

import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)

    // Loop for each test case
    for t > 0 {
        var x, y, a int

        // Input for each test case using fmt.Scanf
        fmt.Scanf("%d %d %d", &x, &y, &a)

        if(a >= x && a < y){
            fmt.Println("yes")
		} else {
			fmt.Println("no")
	
        }

        t--
    }
}
