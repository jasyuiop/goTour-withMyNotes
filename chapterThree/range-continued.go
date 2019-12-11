package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// burada her döngüde i'yi
		// 2**i şeklinde kullanarak
		// pow'u dolduruyorum.
		pow[i] = 1 << uint(i)
	}

	// index veya value yerine "_" karakterini koyarak boş geçebiliyoruz
	/* for i, _ := range pow
	   for _, value := range pow
	*/
	// sadece index'i istiyorsak ikinci değeri atlayıpta yazabiliyoruz
	// for i := range pow

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
