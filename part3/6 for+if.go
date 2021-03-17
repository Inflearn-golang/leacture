package main

import "fmt"

func main() {
	var score, total int
	for v := 0; v < 3000; v++ {
		// total = 0
		for i := 0; i < 3; i++ {
			fmt.Scanln(&score)
			total += score
		}

		if total/3 > 80 {
			fmt.Println("합격")
		} else {
			fmt.Println("불합격")
		}
	}
}
