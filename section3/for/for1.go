package main

import "fmt"

func main() {
	// 반복문 - for
	// Go에서 유일하게 제공되는 반복문
	// 다양한 사용법 숙지하기

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 : ", i)
	}

	// 에러 발생 1
	// for i:=0; i<5; i++
	// {
	// 	fmt.Println("ex 1 : ", i)
	// }

	// 에러 발생 2
	// for i:=0; i<5; i++
	// 	fmt.Println("ex 1 : ", i)

	// 예제 2 (무한 루프)
	// for {
	// 	fmt.Println("ex 2: hello golang")
	// 	fmt.Println("ex 2: infinite loop")
	// }

	// 예제 3 (Range 문법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex 3 :", index, name)
	}
	for _, name := range loc {
		fmt.Println("ex 3 :", name)
	}
}
