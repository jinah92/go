package main

// 선언방법 1
/*
import "fmt"
import "os"
*/

// 선언방법 2
import (
	"fmt"
	"os"
)

func main() {
	// 패키지 : 코드 구조화(모듈) 및 재사용
	// 응집도, 결합도
	// Go : 패키지 단위의 독립이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용함
	// 네이밍 규칙: 소문자(private), 대문자(public)
	// Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 x, 프로그램의 시작점

	var name string

	fmt.Println("이름은? :")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}