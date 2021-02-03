// IF문 (2)
package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// 예제 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 미만")
	}
	
	if b >= 65 {
		fmt.Println("65 이상")
	} else {
			fmt.Println("65 미만")
	}

	// error
	/* if a >= 65{
		fmt.Println("65 이상")
	} else
	{
		fmt.Println("65 미만")
	 } */
}