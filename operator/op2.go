package main

import "fmt"

func main(){
	var a bool
	a = 3 > 4

	fmt.Println(a)

	a = 3 < 4 && 2 < 5
	fmt.Println(a)
}