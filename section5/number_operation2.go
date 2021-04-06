// 데이터 타입 : Numeric 연산(2)

package main

import "fmt"

func main() {
	// 예제 1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n1)
	
	// 예제 2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// fmt.Println("ex3 : ", n3 + n4) // 예외 발생
	fmt.Println("ex3 : ", float32(n3) + n4)
	fmt.Println("ex3 : ", n3 + int(n4))
	fmt.Println("ex3 " , n5 + uint16(n6))
}