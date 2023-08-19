package main

import "fmt"

func main() {
	generateTriangle(5)
	ganerateRect(5)
	
}

func ganerateRect(size int) {
	for i := 1; i < size; i++ {
		for j := 1; j < size-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
			fmt.Println()

	}

}


func generateTriangle(size int) {
	for i := size; i >= 1; i-- {
		for j := 1; j <= size-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
