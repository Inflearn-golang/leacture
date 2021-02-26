package main

import "fmt"

func main() {
	a := 1
	for a <= 2 {
		fmt.Println(a)
	}

	for {
		fmt.Println("안녕하세요 손님.")
		b := 0
		fmt.Scanln(&b)
		fmt.Print(b)
		fmt.Println(" 빠르게 준비해드리겠습니다.")
	}
}
