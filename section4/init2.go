// go 초기화 함수(2)

package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init Method Start 1!")
}
func init() {
	fmt.Println("Init Method Start 2!")
}
func init() {
	fmt.Println("Init Method Start 3!")
}
func init() {
	fmt.Println("Init Method Start 4!")
}

func main() {
	fmt.Println("main start")
}