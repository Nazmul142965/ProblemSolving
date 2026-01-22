package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanln(&x, &y)

	var time int
	time = x * y

	fmt.Println(time)
}
