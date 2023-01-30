package main

import "fmt"

func getArea(resultArea chan int, base int) {
	resultArea <- base * base
}

func main() {
	resultArea := make(chan int, 1)
	go getArea(resultArea, 5)

	fmt.Println(<-resultArea)
}
