package main

import "fmt"

func main() {
	// 제어문 - switch
	// Switch 뒤 표현식(Expression) 생략 가능
	// case 뒤 표현식(Expression) 사용 가능
	// 자동 break 때문에 fallthrouth 존재
	// Type 분기 => 값이 아닌 변수 타입으로 분기 가능

	// 예제 1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제 2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제 3
	switch c := "go"; c {
	case "go":
		fmt.Println("GO!")
	case "java":
		fmt.Println("JAVA!")
	default:
		fmt.Println("일치하는 값 없음.")
	}

	// 예제 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("none")
	}

	// 예제 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
