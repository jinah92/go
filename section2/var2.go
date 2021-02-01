package main

import "fmt"

func main() {
	// 여러 개 선언 (보다 가독성이 좋아짐)
	var (
		name string = "machine"
		height int32
		weight float32
		isRunning bool
	)

	// 변수 초기화
	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name: ", name, "height: ", height, "weight: ", weight, "isRunning: ", isRunning)
}