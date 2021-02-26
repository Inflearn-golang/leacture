package main

import "fmt"

func main() {
	z := []string{"아메리카노", "카페라떼", "돌체라떼"}

	for {
		fmt.Println("안녕하세요 손님.")
		for idx, menu := range z {
			fmt.Println(idx, menu)
		}
		b := 0
		fmt.Println(&b)
		fmt.Print(z[b])
		fmt.Println(" 빠르게 준비해드리겠습니다.")
	}
}
