package main

import "fmt"

// while döngüsünü for söz deyimi ile bu şekilde yazabiliyoruz.
// ayrıyeten dilde while diye bir keyword yok.
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
