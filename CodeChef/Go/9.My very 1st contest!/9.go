package main

import "fmt"

func main() {
	var n, a, b int

	fmt.Scanf("%d %d %d", &n, &a, &b)

	var c int = n - a
	var d int = c - b

	fmt.Println(c, d)
}
