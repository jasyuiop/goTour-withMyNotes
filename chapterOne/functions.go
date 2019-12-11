package main

//package main

import "fmt"

func topla(x, y int) int {
	return x + y
}
func toplaCarp(x, y int) (toplama int, carpma int) { // geriye iki değer de dönderebiliyorum.
	toplama = x + y
	carpma = x * y
	return // bu naked returndur yukarıdaki değerleri otomatik döndürür
	// çok karışık fonksiyonlarda kullanma
}
func yazdir(x, y string) (string, string) {
	return x, y
}
func main() {
	fmt.Println(topla(3, 5))

	fmt.Println(toplaCarp(10, 15))

	fmt.Println(yazdir("emre", "nefesli"))
}
