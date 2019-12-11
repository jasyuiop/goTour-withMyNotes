package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {

	// func append(s []T, vs ...T) []T
	var s []int
	printSlice(s)

	// eklediğin elemanların tipi ekleyeceğin slice'ın tipi aynı olmak zorunda

	// append nil slice'larda çalışabilir
	s = append(s, 0)
	printSlice(s)

	// slice'a yeni eleman eklendikce kapasitesi de büyür.
	s = append(s, 1)
	printSlice(s)

	// append ile birden fazla eleman da eklenebilir
	s = append(s, 4, 5, 6, 7)
	printSlice(s)

}
