package main

import "fmt"

func main() {
	var t int

	fmt.Scanf("%d", &t)

	// Loop for each test case
	for t > 0 {
		var x int
		fmt.Scanf("%d", &x)

		if x >= 2000 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}

		t--
	}
}
