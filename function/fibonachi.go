package main

import "fmt"

func main() {
	res := f(10)
	fmt.Println(res)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
