package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	// bir slice'ın uzunluğuna len() ile kapasitesine cap()
	// ile ulaşabiliriz
}
func main() {

	// slice'lar hem uzunluğa hemde kapasiteye sahiptir.

	// slice'ın uzunluğu içerdiği eleman sayısına eşittir.

	// slice'ın kapasitesi referans aldığı array'in
	// eleman sayısına eşittir

	s := []int{2, 3, 4, 5, 6, 7, 8}
	printSlice(s)

	// 0 uzunluk verdim array'in uzunluğu aynı ama
	s = s[:0]
	printSlice(s)

	// slice'ın uzunluğunu genişlettim
	s = s[:4]
	printSlice(s)

	// ilk 2 indisi bıraktım bu kapasiteyi de düşürdü.
	s = s[2:]
	printSlice(s)

	// bir slice'ın uzunluğunu yeterli kapasiteye sahip olması durumunda
	// yeniden slice'lıyarak uzatabiliriz.
}
