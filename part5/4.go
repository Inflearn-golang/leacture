package main

import "fmt"

type person struct {
	name string
	age  int
	job  string
}

func main() {
	a1 := person{"Austin", 30, "Programmer"}
	a2 := person{"Violet", 28, "Designer"}
	a3 := person{name: "Sopia", age: 30, job: "Producer"}

	fmt.Println(a1.name, a2.age, a3.job)
	fmt.Println(a1.name, a2.name, a3.name)

	a1.age = 31
	fmt.Println(a1)
}
