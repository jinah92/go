// 열거형 3
package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)

	const (
		_ = iota + 0.75 * 2
		DEFAULT
		SILVER
		_  // 생략 가능
		PLATINUM
	)

	fmt.Println(A, C)
	fmt.Println(DEFAULT, SILVER, PLATINUM)
	
}