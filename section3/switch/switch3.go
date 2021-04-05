package main

import "fmt"

func main() {
	// 예제 1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2, 4, 6 중 하나인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // i가 1, 3, 5 중 하나인 경우
		fmt.Println("a -> ", a, "는 홀수")
	default:
		fmt.Println("유효하지 않은 값")
	}

	// 예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c":
		fmt.Println("c")
	}
}
