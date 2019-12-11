package main

import "fmt"

func main() {
	// defer sözdeyimi
	// devamındaki işlevin, işlevleri yerine getirilene kadar bekler, return olduğunda
	// çalışır.
	defer fmt.Println("world")

	fmt.Println("hello")
}
