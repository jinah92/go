package main

import "fmt"

func main() {
	// 예제 1 (루프 레이블로 표기 --> 빠져나갈 반복문명)
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex 2: ", i)
	}

Loop2:
	// 에러 발생(Loop 레이블 밑에 관련없는 소스코드 있으면 컴파일 에러)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex 3 : ", i, j)
		}
	}
}
