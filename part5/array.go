package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	var c [2][3]int
	fmt.Println(c)
}
