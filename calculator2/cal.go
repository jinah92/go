package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString(('\n')) //개행문자 전까지 문자열을 읽음
	line = strings.TrimSpace(line) // 공백 제거

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Println("입력하신 숫자는 %d, %d 입니다.", n1, n2)
	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+"{
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-"{
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "*"{
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}