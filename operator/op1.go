package main

import "fmt"

func main(){
	a := 21
	c := a % 10
	a = a / 10
	d := a % 10

	fmt.Printf("첫 번째 수: %v 두 번째 수: %v\n", c, d)
}