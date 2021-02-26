package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 3
	m["b"] = 7
	delete(m, "a")
	fmt.Println(m)

	m["Americano"] = 3500
	m["Cafe latte"] = 3500

	fmt.Println(m)

	fmt.Println(m["Americano"])
	a, exi := m["Dolce latte"]
	fmt.Println(a, exi)
}
