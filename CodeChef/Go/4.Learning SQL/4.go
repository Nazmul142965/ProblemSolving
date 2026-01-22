package main

import "fmt"

func main() {
	var r, c, e int

	fmt.Scanln(&r, &c, &e)

	var t int = (r + e) * c

	fmt.Println(t)
}
