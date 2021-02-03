// 열거형 2
package main

import "fmt"

func main() {
	const (
		A = iota *10
		B
		C
	)

	fmt.Println(A, B, C)
}