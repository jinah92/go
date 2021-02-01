package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용(전역 x), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용할 경우, 코드 가독성을 높일 수 있음 ==> 추천!

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false
	
	// shortVar1 := 6 재할당 불가 (예외 발생)

	fmt.Println("var1: ", shortVar1, " var2: ", shortVar2, " var3: ", shortVar3)

	// example
	if i:=10; i<11 {	// i가 11이하이면 수행
		fmt.Println("test success!")
	}
	
}