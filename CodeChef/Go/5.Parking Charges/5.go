package main

import "fmt"

func main() {
	var x, y, h int

	fmt.Scanln(&x, &y, &h)

	var t int = x + y*(h-1)

	fmt.Println(t)
}
