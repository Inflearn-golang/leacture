package main

import "fmt"

func main() {
	fmt.Println(add(3, 5))
	first := add(1, 2)
	second := addall(3, 4, 5)
	fmt.Println(first + second)
}

func add(a, b int) int {
	return a + b
}

func addall(li ...int) int {
	result := 0
	for _, n := range li {
		result += n
	}
	return result
}
