package main

import "fmt"

func main() {
	s := make([]int, 3)
	var a []int
	s[0] = 10
	fmt.Println(s)
	fmt.Println(a)
	a = append(a, 10)
	fmt.Println(a)

	// 주석을 지워가며 실행해보시기 바랍니다.
	// s := make([]int,7)
	l := s[0:3]
	fmt.Println(l)
}
