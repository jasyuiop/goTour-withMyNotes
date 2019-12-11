package main

import "fmt"

// const değerler := ile atanmaz, "=" kullanılır.
const Pi = 3.14

func main() {
	const Kisi = "Emre"
	// constlar int, string, bool vb tanımlanabilir.
	const a string = "deneme"
	fmt.Println("happy", Pi, "day")
	fmt.Println("Merhaba", Kisi, a)
}
