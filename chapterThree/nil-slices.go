package main

import "fmt"

func main() {
	// uzunluk vermeyerek hem dizi hem slice oluşturdum
	var s []int

	fmt.Println(s, len(s), cap(s)) // kapasiteyi ve uzunluğu yazdırdım
	if s == nil {
		fmt.Println("nil!")
	}
	// nil slice'lar 0 uzunluk ve kapasiteye sahiptir, altında bir eleman yoktur
}
