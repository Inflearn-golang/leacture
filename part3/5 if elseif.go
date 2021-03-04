package main

import "fmt"

func main() {
	var password = "abc"
	var userInput string
	fmt.Scanln(&userInput)

	if password == userInput {
		fmt.Println("와이파이에 접속되었습니다!")
	} else if password != userInput {
		fmt.Println("접속에 실패했습니다.")
	} else {
		fmt.Println("어떠한 경우도 이 구문을 실행할 수 없습니다.")
	}

	if password == "abc" {
		fmt.Println("와이파이에 접속되었습니다!")
	} else if password == "bcd" {
		fmt.Println("관리자로 접속하셨습니다.")
	} else {
		fmt.Println("접속 실패")
	}

	if pass := "abc"; pass == password {
		fmt.Println("와이파이에 접속되었습니다.")
	}

	switch password {
	case "abc":
		fmt.Println("접속 성공")
	case "bcd":
		fmt.Println("관리자 접속")
	}
}
