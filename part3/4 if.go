package main

import "fmt"

func main() {
	var a int = 6
	if a < 5 {
		fmt.Println(a)
	} else {
		fmt.Println("oh no!")
	}

	var password = "abc"
	var userInput string
	fmt.Scanln(&userInput)

	if password == userInput {
		fmt.Println("와이파이에 접속되었습니다!")
	} else {
		fmt.Println("접속에 실패했습니다.")
	}
}
