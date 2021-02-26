package main

import "fmt"

func main() {
	var a = 5
	b := 4
	a++
	b--
	fmt.Println(a, b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
}
