package main

import "fmt"

type Human struct {
	name     string
	surrname string
	age      int
}

func main() {
	fmt.Println(Human{"emre", "nefesli", 20})
}
