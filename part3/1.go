package main

import "fmt"

func main() {
	for 방문자수 := 1; 방문자수 <= 2; 방문자수++ {
		var 이름 string
		fmt.Scanln(&이름)
		fmt.Println(이름 + "님 환영합니다.")
	}
	fmt.Println("2명이 방문했습니다.")
}
