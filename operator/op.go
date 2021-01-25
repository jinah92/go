package main

import "fmt"

func main(){
	a := 4
	b := 2

	// 선언하는 4가지 방법
	// var c int
	// c = 5
	// var d = 6
	// var e = 3.14
	// var f int = 8

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Printf("result2= %v\n", a^b)
}