// 패키지 접근제어(2)

package main

import (
	"fmt"
	checkUp "section4/lib" 	// checkUP : alias
	_ "section4/lib2"		// _ : 사용하지 않지만 임시로 import (= 빈 식별자)
)

func main() {
	// 패키지 접근제어
	// 별칭 사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUP.CheckNum(11))
}