package main

import "fmt"

func main() {
	fmt.Println("sayılıyor")

	// ertlenen(defer) işlev yığınları bir stack'e pushlanır.
	// kendisinden sonraki işlev bittiğinde stack'in içerisindekiler
	// lifo şeklinde çalışır, last in first out.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("bitti")
}
