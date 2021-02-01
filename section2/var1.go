// 변수 1
package main

import "fmt"


func main() {
	// 초기화
	// 정수 타입 : 0, 실수(소수점) : 0.0, 문자열: "", Boolean: true/false
	// 변수명 : 숫자로 첫글자 x, 대소문자 구분O, 문자/숫자/밑줄/특수기호 사용 가능
	// 변수 및 상수 : 함수 내외 사용이 가능
	
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi, Golang!"
	var k = 4.74 //자동 형변환(선언 동시 초기화)
	var l = "Hi, Seoul!"
	var m = true

	a = 4
	b = "hello go~"
	e = 10
	
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
	
}