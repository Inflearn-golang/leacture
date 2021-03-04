package main

import (
	"errors"
	"fmt"
	"os"
)

func a(b int) (int, error) {
	if b >= 10 {
		return -1, errors.New("don't exists")
	}
	return b, nil
}

func main() {
	fmt.Println(a(10))
	fmt.Println(a(8))

	f, err := os.Create("./defer.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "something")
}
