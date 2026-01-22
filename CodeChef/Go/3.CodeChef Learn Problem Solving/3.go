package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter course number:")
	fmt.Scanln(&n)
	var t int = n * 2

	fmt.Println("Total course:", t)
}
