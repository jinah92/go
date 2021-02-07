package main

import "fmt"

func main() {
	var A [10]int

	for i := 0; i < 10; i++ {
		A[i] = i * i
	}
	fmt.Println(A)
}
