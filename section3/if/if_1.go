// IF문
package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF문 : 반드시 Boolean형으로 검사하기 -> 1 또는 0(go에서는 불가함 : 자동 형변환 불가)
	// 소괄호 사용 X
	
	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("15 이상입니다.")
	}

	// 예제 2
	if b >= 25 {
		fmt.Println("25 이상입니다.")
	}
	
	// 에러 발생1 (조건문 바로 뒤에 { 사용해야 함)
	// if b >= 25
	// {

	// }

	// 에러 발생2 (조건문 다음에 괄호 없이 사용불가)
	// if b >= 24
	// 	fmt.Println("error")

	// 에러 발생3 (1 또는 0 사용불가)
	// if c:=1; c {

	// }

	// 예제 3
	if c:=40; c >= 35 {
		fmt.Println("35이상")
	}

	// error
	// c += 20
}