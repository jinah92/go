package main

import "fmt"

func main(){
	x := 33

	switch x {
	case 31:
		fmt.Println("X = 31")
	case 32:
		fmt.Println("x = 32")
	case 33:
		fmt.Println("x = 33")
	}

	// 조건문으 잆으면 default로 true
	// 한번 true가 있으면 바로 빠져나온다.
	switch {
	case x > 40:
		fmt.Println("x는 40보다 크다.")

	case x < 20:
		fmt.Println("x는 20보다 작다.")
	case x > 30:
		fmt.Println("x는 30보다 크다")
	case x > 20:
		fmt.Println("x는 20보다 크다")
	}
}